use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Request, Response, Server};
use pyo3::{
    prelude::{pyfunction, pymodule, IntoPy, PyModule, PyObject, PyResult, Python},
    wrap_pyfunction,
};
use redis::Client;
use std::convert::Infallible;
use std::net::SocketAddr;
use std::sync::{atomic::AtomicUsize, Arc, Mutex};
use tokio;

mod asyncio;
mod config;
mod conversion;
mod exceptions;
mod lifespan;
mod pool;
mod protocol;
mod runtime;
mod server;
mod types;

/// Connect to a redis server at `address` and use up to `pool_size` connections.
#[pyfunction]
#[text_signature = "(address, pool_size)"]
fn create_pool(address: String, pool_size: u16, pubsub_size: u16) -> PyResult<PyObject> {
    let (fut, res_fut, loop_) = asyncio::create_future()?;

    runtime::RUNTIME.spawn(async move {
        let client = Client::open(address);

        match client {
            Ok(client) => {
                let mut connections = Vec::with_capacity(pool_size as usize);
                for _ in 0..pool_size {
                    let connection = client.get_multiplexed_tokio_connection().await;

                    match connection {
                        Ok(conn) => connections.push(conn),
                        Err(e) => {
                            let _ = asyncio::set_fut_exc(
                                loop_,
                                fut,
                                exceptions::ConnectionError::new_err(format!("{}", e)),
                            );
                            return;
                        }
                    }
                }
                let mut pubsub_connections = Vec::with_capacity(pubsub_size as usize);
                for _ in 0..pubsub_size {
                    let connection = client.get_tokio_connection().await;

                    match connection {
                        Ok(conn) => pubsub_connections.push(conn.into_pubsub()),
                        Err(e) => {
                            let _ = asyncio::set_fut_exc(
                                loop_,
                                fut,
                                exceptions::ConnectionError::new_err(format!("{}", e)),
                            );
                            return;
                        }
                    }
                }

                let pool = pool::ConnectionPool {
                    current: AtomicUsize::new(0),
                    pool: connections,
                    pubsub_pool: Arc::new(Mutex::new(pubsub_connections)),
                    pool_size: pool_size as usize,
                };
                let gil = Python::acquire_gil();
                let py = gil.python();
                let _ = asyncio::set_fut_result(loop_, fut, pool.into_py(py));
            }
            Err(e) => {
                let _ = asyncio::set_fut_exc(
                    loop_,
                    fut,
                    exceptions::ConnectionError::new_err(format!("{}", e)),
                );
            }
        }
    });

    Ok(res_fut)
}

// run, pyasgi入口函数，传入app(字符串，表示app的模块路径)，host，port等，按照给定的protocol启动服务
// 并与app进行交互 app: app.app:app
#[pyfunction]
fn run(
    app: String,
    host: Option<String>,
    port: Option<i16>,
    uds: Option<String>,
    fd: Option<i16>,
    debug: Option<bool>,
) -> PyResult<()> {
    runtime::RUNTIME.block_on(async move {
        let gil = Python::acquire_gil();
        let py = gil.python();

        let split: Vec<&str> = app.split(":").collect();

        match py.import(split[0]) {
            Ok(v) => match v.getattr(split[1]) {
                Ok(appfn) => {
                    let config = config::Config::new();
                    let app = server::Server::new(appfn, config);
                    app.serve(appfn).await;
                }
                Err(e) => {
                    println!("模块中不存在这个app对象");
                }
            },
            Err(e) => {
                println!("不存在这个模块");
            }
        };
    });
    Ok(())
    // let gil = Python::acquire_gil();
    // let py = gil.python();

    // let split: Vec<&str> = app.split(":").collect();
    // let module = py.import(split[0])?;
    // let appfn = module.getattr(split[1])?;
    // let config = config::Config::new();
    // let app = server::Server::new(appfn, config);
    // app.serve(appfn).await;
    // Ok(())
}

async fn hello_world(_req: Request<Body>) -> Result<Response<Body>, Infallible> {
    Ok(Response::new("Hello, World".into()))
}
#[pyfunction]
fn tokio_main() -> PyResult<()> {
    async fn main() {
        // We'll bind to 127.0.0.1:3000
        let addr = SocketAddr::from(([127, 0, 0, 1], 3000));

        // A `Service` is needed for every connection, so this
        // creates one from our `hello_world` function.
        let make_svc = make_service_fn(|_conn| async {
            // service_fn converts our function into a `Service`
            Ok::<_, Infallible>(service_fn(hello_world))
        });

        let server = Server::bind(&addr).serve(make_svc);

        // Run this server for... forever!
        if let Err(e) = server.await {
            eprintln!("server error: {}", e);
        }
    }
    runtime::RUNTIME.block_on(main());
    Ok(())
}

#[pymodule]
fn pyasgi(py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(create_pool, m)?)?;
    m.add_function(wrap_pyfunction!(run, m)?)?;
    m.add_class::<pool::ConnectionPool>()?;
    m.add(
        "ConnectionError",
        py.get_type::<exceptions::ConnectionError>(),
    )?;
    m.add("ArgumentError", py.get_type::<exceptions::ArgumentError>())?;
    m.add("RedisError", py.get_type::<exceptions::RedisError>())?;
    m.add("PoolEmpty", py.get_type::<exceptions::PoolEmpty>())?;
    m.add("PubSubClosed", py.get_type::<exceptions::PubSubClosed>())?;

    Ok(())
}
