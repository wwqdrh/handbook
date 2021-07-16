#[macro_use]
extern crate log;
#[macro_use]
extern crate serde_json;
extern crate env_logger;
extern crate url;

use std::collections::HashMap;
use std::io;

use futures::future::{Future, FutureResult};
use futures::Stream;
use hyper::server::{Request, Response, Service};
use hyper::header::{ContentLength, ContentType};
use hyper::Method::{Get, Post};
use hyper::{Chunk, StatusCode};

struct TimeRange {
    before: Option<i64>,
    after: Option<i64>,
}

#[derive(Debug)]
pub struct Message {
    pub id: i32,
    pub username: String,
    pub message: String,
    pub timestamp: i64,
}

#[derive(Debug)]
pub struct NewMessage {
    pub username: String,
    pub message: String,
}

struct Microservice;

impl Service for Microservice {
    type Request = Request;
    type Response = Response;
    type Error = hyper::Error;
    type Future = Box<dyn Future<Item = Self::Response, Error = Self::Error>>;

    fn call(&self, request: Request) -> Self::Future {
        match (request.method(), request.path()) {
            (&Get, "/") => {
                let time_range = match request.query() {
                    Some(query) => parse_query(query),
                    None => Ok(TimeRange {
                        before: None,
                        after: None,
                    }),
                };
                let response = match time_range {
                    Ok(time_range) => make_get_response(query_db(time_range)),
                    Err(error) => make_error_response(&error),
                };
                Box::new(response)
            }
            (&Post, "/") => {
                let future = request
                    .body()
                    .concat2()
                    .and_then(parse_form)
                    .and_then(write_to_db)
                    .then(make_post_response);
                Box::new(future)
            }
            _ => Box::new(futures::future::ok(
                Response::new().with_status(StatusCode::NotFound),
            )),
        }

        info!("Microservice received a request: {:?}", request);
        Box::new(futures::future::ok(Response::new()))
    }
}

fn parse_query(query: &str) -> Result<TimeRange, String> {
    let args = url::form_urlencoded::parse(&query.as_bytes())
        .into_owned()
        .collect::<HashMap<String, String>>();

    let before = args.get("before").map(|value| value.parse::<i64>());
    if let Some(ref result) = before {
        if let Err(ref error) = *result {
            return Err(format!("Error parsing 'before': {}", error));
        }
    }

    let after = args.get("after").map(|value| value.parse::<i64>());
    if let Some(ref result) = after {
        if let Err(ref error) = *result {
            return Err(format!("Error parsing 'after': {}", error));
        }
    }

    Ok(TimeRange {
        before: before.map(|b| b.unwrap()),
        after: after.map(|a| a.unwrap()),
    })
}

fn parse_form(form_chunk: Chunk) -> FutureResult<NewMessage, hyper::Error> {
    let mut form = url::form_urlencoded::parse(form_chunk.as_ref())
        .into_owned()
        .collect::<HashMap<String, String>>();
    if let Some(message) = form.remove("message") {
        let username = form.remove("username").unwrap_or(String::from("anonymous"));
        futures::future::ok(NewMessage {
            username: username,
            message: message,
        })
    } else {
        futures::future::err(hyper::Error::from(io::Error::new(
            io::ErrorKind::InvalidInput,
            "Missing field 'message'",
        )))
    }
}

fn write_to_db(entry: NewMessage) -> FutureResult<i64, hyper::Error> {
    futures::future::ok(0)
}

fn query_db(time_range: TimeRange) -> Option<Vec<Message>> {
    let mut vec: Vec<Message> = Vec::new();
    vec.append(&mut Message {
        id: 1,
        username: "user1".to_string(),
        message: "message1".to_string(),
        timestamp: 1620783274,
    });
    vec.append(&mut Message {
        id: 2,
        username: "user2".to_string(),
        message: "message2".to_string(),
        timestamp: 1620783274,
    });
    return Some(vec);
}

fn make_get_response(
    messages: Option<Vec<Message>>,
) -> FutureResult<hyper::Response, hyper::Error> {
    let response = match messages {
        Some(messages) => {
            let body = render_page(messages);
            Response::new()
                .with_header(ContentLength(body.len() as u64))
                .with_body(body)
        }
        None => Response::new().with_status(StatusCode::InternalServerError),
    };
    debug!("{:?}", response);
    futures::future::ok(response);
}

fn make_post_response(
    result: Result<i64, hyper::Error>,
) -> FutureResult<hyper::Response, hyper::Error> {
    match result {
        Ok(timestamp) => {
            let payload = json!({ "timestamp": timestamp }).to_string();
            let response = Response::new()
                .with_header(ContentLength(payload.len() as u64))
                .with_header(ContentType::json())
                .with_body(payload);
            debug!("{:?}", response)
        }
        Err(error) => make_error_response(error.description()),
    }
}

fn make_error_response(error_message: &str) -> FutureResult<hyper::Response, hyper::Error> {
    let payload = json!({ "error": error_message }).to_string();
    let response = Response::new()
        .with_status(StatusCode::InternalServerError)
        .with_header(ContentLength(payload.len() as u64))
        .with_header(ContentType::json())
        .with_body(payload);
    debug!("{:?}", response);
    futures::future::ok(response)
}

fn render_page(messages: Vec<Message>) -> String {
    return String::from("a html message");
    // (html! {
    //   head {
    //     title "microservice"
    //     style "body { font-family: monospace }"
    //   }
    //   body {
    //     ul {
    //       @for message in &messages {
    //         li {
    //           (message.username) " (" (message.timestamp) "): " (message.message)
    //         }
    //       }
    //     }
    //   }
    // }).into_string()
}

fn main() {
    env_logger::init();
    let address = "127.0.0.1:8080".parse().unwrap();
    let server = hyper::server::Http::new()
        .bind(&address, || Ok(Microservice {}))
        .unwrap();
    info!("Running microservice at {}", address);
    server.run().unwrap();
}
