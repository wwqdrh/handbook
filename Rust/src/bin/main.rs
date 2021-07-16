// mod crate::cook;
// mod data;
// mod server;
// use cook::algo;
// use data::base;
// use data::concurrent;
// use data::feature;
// use data::structure;
// use super::server::http;
use cookbook::server::http;
use cookbook::server::rocket;

fn main() {
    //  * 1、cargo项目管理工具，直接在cargo.toml中添加依赖然后cargo更新
    //  * 2、没有分号就是返回值表达式，直接返回
    // base::int();
    // println!("result, {}", base::int_add(1, 5));
    // base::do_slice();
    // base::do_struct();
    // base::do_enum();

    // structure::do_loop();
    // structure::do_error_handler();

    // feature::do_generic();
    // feature::do_trait();

    // concurrent::spawn_function();
    // concurrent::do_chan();
    // concurrent::do_multi_chan();
    // concurrent::do_mux();

    // algo::rand_num();

    // http::single_http_server();
    // http::thread_pool_http_server();
    rocket::run_rocket_server();
}
