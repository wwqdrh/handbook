use pyo3::prelude::*;

pub struct Config {
    pub host: String,
    pub port: i16,
}

impl Config {
    pub fn new() -> Config {
        Self {
            host: String::from("127.0.0.1"),
            port: 8000,
        }
    }
}
