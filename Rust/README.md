Rust 是现有系统软件语言（如 C 和 C++）的一种安全替代语言。 与 C 和 C++ 一样，Rust 没有大型运行时或垃圾回收器，这几乎与所有其他现代语言形成了鲜明对比。 但是，与 C 和 C++ 不同的是，Rust 保证了内存安全。 Rust 可以避免很多与在 C 和 C++ 中遇到的内存使用错误相关的 bug。

类型安全、内存安全、无数据争用、零成本抽象、最小运行时、面向裸机

## 模块系统

包：
包含一个或多个 crate 内的功能。
包括有关如何生成这些 crate 的信息。 该信息位于 Cargo.toml 文件中。

箱：
是编译单元，即 Rust 编译器可以运行的最小代码量。
编译完成后，系统将生成可执行文件或库文件。
其中包含未命名的隐式顶层模块。

模块：
是箱内的代码组织单位（或为嵌套形式）。
可以具有跨其他模块的递归定义。

``` Shell
my-project
├── src
│  └── main.rs
└── Cargo.toml
```

src/bin目录可以放置多个文件，包中可以包含多个二进制箱，如果包中包含 src/main.rs 和 src/lib.rs，则其中有两个箱：库文件和二进制文件。

`cargo new --lib my-library` 是生成库的命令，可以链接到其他程序。

如果源文件中存在 mod 声明，则在运行编译器之前，系统会将模块文件的内容插入到 mod 声明在源文件中的所在位置。 换句话说，系统不会对模块进行单独编译，只会编译箱。

模块中的项如果声明为pub，就视为可供外界访问。

### 常用crate

std::collections - 集合类型的定义，如 HashMap。
std::env - 用于处理环境的函数。
std::fmt - 控制输出格式的功能。
std::fs - 用于处理文件系统的功能。
std::io - 用于处理输入/输出的定义和功能。
std::path - 支持处理文件系统路径数据的定义和功能。
structopt - 用于轻松分析命令行参数的第三方 crate。
chrono - 用于处理日期和时间数据的第三方箱。
regex - 用于处理正则表达式的第三方箱。
serde - 适用于 Rust 数据结构的序列化和反序列化操作的第三方箱。

### cargo

用cargo来做依赖项管理

使用 cargo new 命令创建新的项目模板。
使用 cargo build 编译项目。
使用 cargo run 命令编译并运行项目。
使用 cargo test 命令测试项目。
使用 cargo check 命令检查项目类型。
使用 cargo doc 命令编译项目的文档。
使用 cargo publish 命令将库发布到 crates.io。
通过将箱的名称添加到 Cargo.toml 文件来将依赖箱添加到项目。(只能手动编辑文件)

## 所有权

所有权和借用，使得Rust无需垃圾回收器就能实现内存安全保证

当对象超出作用域(一般是{}代码块内的)，会被删除，内存将释放。

可以对所有权进行转移，不过一旦移动，旧变量将不再有效

## 测试

使用`#[test]`标记的简单函数，通常使用 assert! 或 assert_eq! 宏来检查结果。

一般与源代码放在一起，然后使用cargo test命令执行测试输出

还有

- 文档测试: 在代码注释块中的测试文件也可以执行
- 集成测试: 将crate作为整体进行测试，tests目录下，cargo会运行此目录中的每个源文件

``` Rust
fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod add_function_tests {
    use super::*;

    #[test]
    fn add_works() {
        assert_eq!(add(1, 2), 3);
        assert_eq!(add(10, 12), 22);
        assert_eq!(add(5, -2), 3);
    }

    #[test]
    #[should_panic]
    fn add_fails() {
        assert_eq!(add(2, 2), 7);
    }

    #[test]
    #[ignore]
    fn add_negatives() {
        assert_eq!(add(-2, -2), -4)
    }
}
```