use std::collections::HashMap;

pub enum LifespanReceiveMessage {
    startup,
    shutdown,
}

impl LifespanReceiveMessage {
    pub fn to_hashmap(&self) -> HashMap<String, String> {
        match &self {
            Self::startup => {
                let mut res: HashMap<String, String> = HashMap::new();
                res.insert(String::from("type"), String::from("lifespan.startup"));
                return res;
            }
            Self::shutdown => {
                let mut res: HashMap<String, String> = HashMap::new();
                res.insert(String::from("type"), String::from("lifespan.startdown"));
                return res;
            }
        }
    }
}

pub struct ASGISpecInfo {
    pub version: String,
    pub spec_version: String, // 2.0, 2.1
}
pub struct LifespanScope {
    pub type_: String,
    pub asgi: ASGISpecInfo,
}
