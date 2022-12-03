```bash
$ cargo build --target wasm32-wasi --release

$ wasmedge ./target/wasm32-wasi/release/wasi_hello_world.wasm

thread 'main' panicked at 'called `Result::unwrap()` on an `Err` value: Custom { kind: Uncategorized, error: "failed to find a pre-opened file descriptor through which \"./helloworld.txt\" could be opened" }', src/main.rs:18:18
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
[2022-12-03 15:08:43.800] [error] execution failed: unreachable, Code: 0x89
[2022-12-03 15:08:43.800] [error]     In instruction: unreachable (0x00) , Bytecode offset: 0x00006d86
[2022-12-03 15:08:43.800] [error]     When executing function name: "_start"
```
wasm运行时，对于文件、网络等需要开启相应权限

对于文件就是挂载文件路径，`--dir guest:host`

```bash
$ wasmedge --dir /helloworld:./temp ./target/wasm32-wasi/release/wasi_hello_world.wasm
```