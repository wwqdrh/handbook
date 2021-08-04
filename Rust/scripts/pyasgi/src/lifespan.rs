// 定义asgi的lifespan行为
use std::collections::{HashMap, VecDeque};
use std::sync::Arc;

use crate::{config::Config, runtime};
use crate::types::LifespanReceiveMessage;

pub struct LifeSpanOn<'a> {
    config: &'a Config,
    receive_queue: &'a VecDeque<HashMap<String, String>>,
    should_exit: &'a bool,
}

impl<'a> LifeSpanOn<'a> {
    pub fn new(config: Config) -> LifeSpanOn<'a> {
        Self {
            config: config,
            receive_queue: VecDeque::new(),
            should_exit: false,
        }
    }
    // lifespan的主循环
    async fn main(&self) {
        // let app = &self.config.app;
        // 给app发送消息接收消息等等
    }

    pub async fn serve(&self) {}

    pub async fn startup(&self) {
        println!("Waiting for application startup.");
        let lifespan = Arc::new(self);
        runtime::RUNTIME.spawn(async move {
            lifespan.main().await;
        });
        self.receive_queue.push_back(LifespanReceiveMessage::startup.to_hashmap());
        
    }

    pub async fn main_loop(&self) {}

    pub async fn shutdown(&self) {}
}

pub struct LifeSpanOff {}

impl LifeSpanOff {}

