use crate::lifespan::LifeSpanOn;
use crate::protocol::http;
use crate::{
    asyncio::{self, get_loop},
    config::Config,
    runtime,
};
use futures_util::TryStreamExt;
use hyper::service::{make_service_fn, service_fn};
use pyo3::prelude::*;
use pyo3::{PyAny, PyResult};
// use hyper::{Body, Method, Request, Response, Server, StatusCode};
// use crate::lifespan::LifeSpan;
use std::collections::HashSet;
struct ServerState {
    total_requests: i16,
    connections: HashSet<String>,
    tasks: HashSet<String>,
    default_headers: Vec<String>,
}

pub struct Server {
    config: Config,
    server_state: ServerState,
    started: bool,
    should_exit: bool,
    lifespan: LifeSpanOn,
}

impl Server {
    pub fn new(app: &PyAny, config: Config) -> Server {
        Self {
            config: config,
            server_state: ServerState {
                total_requests: 0,
                connections: HashSet::new(),
                tasks: HashSet::new(),
                default_headers: vec![],
            },
            started: false,
            should_exit: false,
            lifespan: LifeSpanOn::new(config),
        }
    }

    pub async fn serve(&self, app: &PyAny) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
        println!(
            "Started server process \
            Started server process
        "
        );

        let gil = Python::acquire_gil();
        let py = gil.python();

        self.startup().await;
        self.main_loop().await;
        self.shutdown().await;
        // asyncio::run_async(app.call1((1.into_py(py), 2.into_py(py), 3.into_py(py)))?);

        // loop_.call_method1(py, "create_task", (app, 1, 2, 3))?;
        // runtime::RUNTIME.block_on(app.call1((1, 1, 1))?);
        // self.startup().await;
        // self.main_loop().await;
        // self.shutdown().await;

        println!(
            "Finished server process \
            Finisahed server process
        "
        );

        Ok(())
    }

    async fn startup(&'static self) {
        self.lifespan.startup().await;
        // 创建server
        let addr = ([127, 0, 0, 1], 3000).into();

        let service = make_service_fn(|_| async { Ok::<_, hyper::Error>(service_fn(http::echo)) });

        let server = hyper::Server::bind(&addr).serve(service);

        println!("Listening on http://{}", addr);

        server.await?;
    }

    async fn main_loop(&self) {}

    async fn shutdown(&self) {}
}
