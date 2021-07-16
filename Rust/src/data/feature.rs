pub fn do_generic() {
    // 结构体与枚举类都可以定义方法，并且能够实现泛型

    struct Point<T> {
        x: T,
        y: T,
    }

    impl<T> Point<T> {
        fn x(&self) -> &T {
            &self.x
        }
    }

    let p = Point { x: 1, y: 2 };
    println!("p.x = {}", p.x());
}

pub fn do_trait() {
    // trait，类似于Java中的接口，都是一种行为规范
    // trait可以定义方法作为默认方法，该方法也支持实现结构体重新定义
    // trait可以作为函数的参数进行传递
    // trait, impl for
    // 继承是多态思想(可以处理多种类型数据的代码)的实现，在rust中
    // 通过trait实现多态，但是无法实现属性的继承，想继承一个类的方法最好在子类中定义父类的实例，没有语法糖也没有官方继承手段

    trait Descriptive {
        fn describe(&self) -> String;
    }

    struct Person {
        name: String,
        age: u8,
    }

    impl Descriptive for Person {
        fn describe(&self) -> String {
            format!("{} {}", self.name, self.age)
        }
    }

    fn output(obj: impl Descriptive) {
        println!("{}", obj.describe());
    }
}
