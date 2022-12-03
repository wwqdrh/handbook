use std::{fs, io::prelude::*};

/**
 * webserver主要涉及两个主要协议、超文本传输(http)以及传输控制协议(tcp)
*/
use std::{
    net::{TcpListener, TcpStream},
    thread,
};
use crate::server::pool::ThreadPool;

fn handle_connection(mut stream: TcpStream) {
    let mut buffer = [0; 1024]; // 1024个字节的缓冲区
    stream.read(&mut buffer).unwrap();
    // println!("Request: {}", String::from_utf8_lossy(&buffer[..]));

    let get = b"GET / HTTP/1.1\r\n";

    if buffer.starts_with(get) {
        let contents = fs::read_to_string("static/server/index.html").unwrap();

        let response = format!(
            "HTTP/1.1 200 OK\r\nContent-Length: {}\r\n\r\n{}",
            contents.len(),
            contents
        );
        // println!("{}", response);
        stream.write(response.as_bytes()).unwrap();
        stream.flush().unwrap();
    } else {
        let contents = fs::read_to_string("static/server/404.html").unwrap();

        let response = format!(
            "HTTP/1.1 404 NOT FOUND\r\nContent-Length: {}\r\n\r\n{}",
            contents.len(),
            contents
        );
        // println!("{}", response);
        stream.write(response.as_bytes()).unwrap();
        stream.flush().unwrap();
    }
}

pub fn single_http_server() {
    // 单线程httpserver
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();
    println!(":listen[7878]");
    for stream in listener.incoming() {
        let stream = stream.unwrap();
        println!("Connection established!");
        handle_connection(stream);
    }
}

pub fn each_thread_http_server() {
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();

    for stream in listener.incoming() {
        let stream = stream.unwrap();

        thread::spawn(|| {
            handle_connection(stream);
        });
    }
}

pub fn thread_pool_http_server() {
    let listener = TcpListener::bind("127.0.0.1:7878").unwrap();
    let pool = ThreadPool::new(4);

    for stream in listener.incoming().take(2) {
        let stream = stream.unwrap();

        pool.execute(|| {
            handle_connection(stream);
        });
    }

    println!("Shutting down.");
}
