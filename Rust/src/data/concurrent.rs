/**
 * 如何安全的共享内存，rust提供了两种方法，通过通道来共享内存，或者使用互斥器来保证共享内存的安全性
 * 使用std::marker::{Sync Send} 的trait实现自己的可扩展并发
 * - 通过Send允许在线程间转移所有权
 * Send 标记 trait 表明类型的所有权可以在线程间传递。
 * 几乎所有的 Rust 类型都是Send 的，
 * 不过有一些例外，包括 Rc<T>：这是不能 Send 的，
 * 因为如果克隆了 Rc<T> 的值并尝试将克隆的所有权转移到另一个线程，这两个线程都可能同时更新引用计数。
 * 为此，Rc<T> 被实现为用于单线程场景，这时不需要为拥有线程安全的引用计数而付出性能代价。
 * 任何完全由 Send 的类型组成的类型也会自动被标记为 Send。几乎所有基本类型都是 Send 的
 * - Sync允许多线程访问
 * Sync 标记 trait 表明一个实现了 Sync 的类型可以安全的在多个线程中拥有其值的引用。
 * 换一种方式来说，对于任意类型 T，如果 &T（T 的引用）是 Send 的话 T 就是 Sync 的，这意味着其引用就可以安全的发送到另一个线程。
 * 类似于 Send 的情况，基本类型是 Sync 的，完全由 Sync 的类型组成的类型也是 Sync 的。
*/
use std::sync::mpsc;
use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Duration;

pub fn spawn_function() {
    //  * 线程（thread）是一个程序中独立运行的一个部分。
    // 线程不同于进程（process）的地方是线程是程序以内的概念，程序往往是在一个进程中执行的。
    // 在有操作系统的环境中进程往往被交替地调度得以执行，线程则在进程以内由程序进行调度。
    // 由于线程并发很有可能出现并行的情况，所以在并行中可能遇到的死锁、延宕错误常出现于含有并发机制的程序。
    // 为了解决这些问题，很多其它语言（如 Java、C#）采用特殊的运行时（runtime）软件来协调资源，但这样无疑极大地降低了程序的执行效率。
    // C/C++ 语言在操作系统的最底层也支持多线程，且语言本身以及其编译器不具备侦察和避免并行错误的能力，这对于开发者来说压力很大，开发者需要花费大量的精力避免发生错误。
    // Rust 不依靠运行时环境，这一点像 C/C++ 一样。
    // 但 Rust 在语言本身就设计了包括所有权机制在内的手段来尽可能地把最常见的错误消灭在编译阶段，这一点其他语言不具备
    // spawn方法创建线程，join线程等待，move强制所有权迁移，消息传递
    let handle = thread::spawn(|| {
        for i in 0..5 {
            println!("spawned thread print {}", i);
            thread::sleep(Duration::from_millis(1));
        }
    });
    handle.join().unwrap();

    // 当子线程中尝试使用当前函数的资源，会发出错误，因为所有权机制禁止这种危险情况的产生，
    // 它将破会所有权机制销毁资源的一定性，使用闭包的move关键字来处理
    let s = "hello";
    let handle = thread::spawn(move || {
        println!("{}", s);
    });
    handle.join().unwrap();
}

pub fn do_chan() {
    // 消息传递，一个实现消息传递的主要工具是channel，一个发送者一个接受者
    // std::sync::mpsc
    let (tx, rx) = mpsc::channel();
    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
        ];
        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        // 当通道关闭的时候迭代器也会关闭
        println!("Got: {}", received);
    }
}

pub fn do_multi_chan() {
    let (tx, rx) = mpsc::channel();
    let tx1 = mpsc::Sender::clone(&tx);

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
        ];
        for val in vals {
            tx.send(val).unwrap();
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("more"),
            String::from("message"),
            String::from("for"),
        ];
        for val in vals {
            tx1.send(val).unwrap()
        }
    });

    for received in rx {
        println!("Got: {}", received);
    }
}

pub fn do_mux() {
    // new 创建锁，lock获取锁，如果锁被其他线程拥有了并且那个线程panic了，则lock调用会失败
    // let m = Mutex::new(5);

    // {
    //     let mut num = m.lock().unwrap(); // 离开作用域后会自动释放锁
    //     *num = 6;
    // }

    // println!("m = {:?}", m);

    // 需要注意的是不能被多次移动，可以通过多所有权来处理,
    // std::rc::Rc, 并非线程安全，没有任何并发原语，因此可能操作会被其他线程打断， 或者计数错误导致诡异的bug
    // std::sync::{Mutex, Arc}, 线程安全

    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();

            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());
}
