/**
 * 整数型: i8 u8 i16 u16 i32 u32 i64 u64
 * 浮点型: f32 f64
 * 布尔型: bool
 * 字符型: char(4个字节，代表Unicode标量值)
 * 复合类型: 元组() 可以包含不同类型, 数组[] 只能包含同种类型
 * 字符串类型: str String, str是Rust核心语言类型，字符串切片，常常以引用的形式出现，用双引号包含的的类型性质都是&str
 *  String是标准库提供的数据类型，支持字符串的追加、清空等操作，都支持切片，切片结果就是&str类型的数据
 *  切片结果必须是引用类型，但开发者必须自己明示这一点 let slice = &s[0..3];
 *  除了字符串，其他一些线性数据结构也支持切片操作
*/

pub fn int() {
    let a = 5;
    let b = 10;

    println!("restult: {}", a + b);
}

pub fn int_add(a: i32, b: i32) -> i32 {
    return a + b;
}

pub fn do_slice() {
    // ..y == 0..y; x.. == x到结束; .. == 从0到结束
    // 尽量不要在字符串中使用费英文字符，被切片引用的字符串禁止更改值
    let s = String::from("rustlearn");

    let part1 = &s[0..4];
    let part2 = &s[4..9];
    println!("{}={}+{}", s, part1, part2);
}

pub fn do_struct() {
    /**
     * 结构体、元组结构体，单元结构体
     * 结构体必须掌握字段值所有权，因为结构体失效的时候会释放所有字段，
     * 例如使用String而不是&str，结构体可以通过生命周期机制来实现定义引用型字段
     */

    // 声明结构体
    #[derive(Debug)] // 导入调试库，这样可以使用{:?}占位输出一整个结构体 {:#?}属性较多可以使用这个
    struct Site {
        domain: String,
        name: String,
        nation: String,
    }

    // 声明元组结构体，更简单的声明方式
    struct Color(u8, u8, u8);

    let black = Color(0, 0, 0);
    // 元组结构体对象的使用方式和元组一样，通过.和下标类进行访问
    println!("black = ({}, {}, {})", black.0, black.1, black.2);

    // 结构体实例
    let rust = Site {
        domain: String::from("www.rust.com"),
        name: String::from("this is rust web"),
        nation: String::from("China"),
    };

    // 根据现有的实例属性新建一个新的 ..实例, 这里rust的所有权被借了，上面的
    // rust实例无法使用了
    let rust2 = Site {
        name: String::from("rust2 web"),
        ..rust
    };

    // 输出结构体 :? 符号占位输出
    println!("rect1 is {:?}", rust2);

    // 定义结构体方法，可以理解为实例方法，类似go语言的机制, impl关键字，内部的方法第一个参数需要是&self
    // 定义结构体函数，可以理解为静态函数，没有&self参数
    // 结构体impl可以写多次，效果相当于内容的拼接
    impl Site {
        fn to_str(&self) -> String {
            return String::from(&self.name);
        }
        fn factory(name: String, domain: String, nation: String) -> Site {
            Site {
                name: name,
                domain: domain,
                nation: nation,
            }
        }
    }

    // 调用结构体的方法
    let rust3 = Site::factory(
        String::from("hhh"),
        String::from("domain"),
        String::from("zh"),
    );
    println!("调用结构体方法{}", rust3.to_str());
    println!("rust3 {:?}", rust3);
}

pub fn do_enum() {
    /**
     * match语法自动避免由于忘了break而导致的串接运行
     *  match 枚举类实例 {
            分类1 => 返回值表达式,
            分类2 => 返回值表达式,
            ...
        }
    * Option是标准库中的枚举类，用于填补Rust不支持null引用的空白
    * null经常性的导致一些bug的出现 Rust 在语言层面彻底不允许空值 null 的存在，但无奈null 可以高效地解决少量的问题，所以 Rust 引入了 Option 枚举类：
    *
    * 使用if let简化match处理 if let 匹配值 = 原变量 {}
    */
    // 枚举类，rust中的枚举类相较于其他语言中更复杂
    // 常规定义
    #[derive(Debug)]
    enum Book {
        Papery, // 还可以加上类型Papery(u32)
        Electronic,
    }

    let book = Book::Papery;
    println!("{:?}", book);

    // 为属性命名，不能直接访问属性，只能通过match访问
    #[derive(Debug)]
    enum Book2 {
        Papery { index: u32 }, // 还可以命名 Papery {index: u32}
        Electronic { url: String },
    }
    let book = Book2::Papery { index: 1001 };
    // let ebook = Book2::Electronic {
    //     url: String::from("url://"),
    // };
    match book {
        Book2::Papery { index } => {
            println!("Papery book {}", index);
        }
        Book2::Electronic { url } => {
            println!("E-book {}", url);
        }
    }

    // 定义一个可以为空值的类
    let opt = Option::Some("Hello");
    // 初始值为空的Option必须明确类型
    // let opt: Option<&str> = Option::None;
    match opt {
        Option::Some(val) => {
            // 也可以省略直接写Some()=>...
            println!("{}", val);
        }
        Option::None => {
            println!("opt is nothing");
        }
    }

    // 使用if let简化match处理
    enum Book3 {
        Papery(u32),
        Electronic(String),
    }
    let book3 = Book3::Electronic(String::from("http://url"));
    if let Book3::Electronic(url) = book3 {
        println!("Electronic url {}", url);
    } else {
        println!("no electronic book");
    }
}
