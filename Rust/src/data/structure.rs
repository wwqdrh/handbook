pub fn do_loop() {
    let mut number = 1;
    while number != 4 {
        println!("{}", number);
        number += 1;
    }
    println!("while loop done");

    let a = [10, 20, 30, 40, 50];
    for i in a.iter() {
        println!("值为: {}", i)
    }
    println!("for loop done");

    // loop循环是无限循环，可以使用 break i的语法类似return一样返回数据
    let s = ['R', 'u', 's', 't'];
    let mut i = 0;
    let location = loop {
        let ch = s[i];
        if ch == 'u' {
            break i;
        }
        println!("\'{}\'", ch);
        i += 1
    };
    println!("\'u\'的索引为{}", location);
}

pub fn do_error_handler() {
    /**
     * 执行异常处理, 分为不可恢复错误panic!与可恢复错误Result<T, E>
     * 可恢复错误包括: 文件访问失败
     *
     * 直接变为panic: 如果想使一个可恢复错误按不可恢复错误处理，Result提供了两个办法，unwrap(), expect(message: &str);
     * 返回给上一层: ? 符的实际作用是将 Result 类非异常的值直接取出，如果有异常就将异常 Result 返回出去。
     *
     * 前面是接收错误，自己的函数传递错误，可以使用Ok, Err, 函数返回类型Result<typea, typeb>
     *
     * kind方法，类似try快一样可以领任何位置发生异常直接得到相同的解决的语法
     * 我们完全可以把 try 块在独立的函数中实现，将所有的异常都传递出去解决。实际上这才是一个分化良好的程序应当遵循的编程方法：应该注重独立功能的完整性。
     * 这样需要判断Result的Err类型，获取Err类型的函数是kind
     */
    use std::fs::File;
    // panic!("error occured");
    let f = File::open("hello.txt");
    if let Ok(file) = f {
        println!("打开文件成功");
    } else {
        println!("failed to open the file");
    }

    // let f1 = File::open("hello.txt").unwrap(); // 如果处理成功会直接把结果返回出来，如果不能直接panic
    // let f2 = File::open("hello.txt").expect("Failed to open."); // 如果没有成功将自定义的message传入panic

    fn action(i: i32) -> Result<i32, bool> {
        if i >= 0 {
            Ok(i)
        } else {
            Err(false)
        }
    }

    let r = action(1000);
    if let Ok(v) = r {
        println!("Ok: {}", v);
    } else {
        println!("Err");
    }

    // kind方法
    use std::io;
    use std::io::Read;

    fn mutil_err(path: &str) -> Result<String, io::Error> {
        let mut f = File::open(path)?;
        let mut s = String::new();
        f.read_to_string(&mut s)?;
        Ok(s)
    }

    let str_file = mutil_err("hello.txt");
    match str_file {
        Ok(s) => println!("{}", s),
        Err(e) => match e.kind() {
            io::ErrorKind::NotFound => {
                println!("No such file");
            }
            _ => {
                println!("Cannot read the file");
            }
        },
    }
    println!("errhandler done");
}
