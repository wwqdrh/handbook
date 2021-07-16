#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use]
extern crate rocket;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

pub fn run_rocket_server() {
    rocket::ignite().mount("/", routes![index]).launch();
}
