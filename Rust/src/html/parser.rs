use std::collections::HashMap;

use nipper::{Document, Selection};

use anyhow::Result;
use lazy_static::lazy_static;
use regex::Regex;


#[derive(Clone, Debug)]
pub struct Info {
    pub max_volume: usize,
    pub project_count: usize,
    pub star: String,
}

pub fn request() -> Info {
    let resp = reqwest::blocking::get("https://hellogithub.com").unwrap();
    parse_hg_info(resp.text().unwrap())
}

/// 返回最大期数
pub fn parse_hg_info(html: String) -> Info {
    let doc = Document::from(&html);

    let text = doc.select(
        "body > div.l-content > div.pricing-tables.pure-g > div:nth-child(2) > div > div > span",
    )
    .text();
    let result = text.trim().split(' ').into_iter().collect::<Vec<&str>>();
    let project_count = result.get(0).unwrap().parse().unwrap();

    let text = doc.select("body > div.l-content > div.pricing-tables.pure-g > div:nth-child(1) > div > div > span").text();
    let result = text.trim().split(' ').into_iter().collect::<Vec<&str>>();
    let max_volume = result.get(0).unwrap().parse().unwrap();

    Info {
        max_volume,
        project_count,
        star: "55.2k".to_string(),
    }
}

mod test {
    #[allow(unused_imports)]
    use super::*;

    #[test]
    #[ignore]
    fn test_volume() {
        let info = request();
        println!("start{:?}", info.star);
    }
}