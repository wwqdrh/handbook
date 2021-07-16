/**
 * 包含rust的高级特性，比如
 * - 解引用裸指针
 * - 调用不安全的函数或方法
 * - 访问或修改可变静态变量
 * - 实现不安全的trait
 * - 访问union的字段
 *
 * 指针：引用、智能指针、裸指针
 * 裸指针：
 * 允许忽略借用规则，可以同时拥有不可变和可变的指针，或多个指向相同位置的可变指针
 * 不保证指向有效的内存
 * 允许为空
 * 不能实现任何自动清理功能
*/

pub fn do_raw_pointers() {
    // 允许忽略借用规则，可以同时拥有不可变和可变的指针，或多个指向相同位置的可变指针
    // 不保证指向有效的内存
    // 允许为空
    // 不能实现任何自动清理功能

    let mut num = 5;

    // 可以在安全代码中创建裸指针，不能再不安全块之外解引用裸指针, 要解引用需要unsafe关键字
    let r1 = &num as *const i32; // 不可变裸指针
    let r2 = &mut num as *mut i32; // 可变裸指针

    unsafe {
        println!("r1 is: {}", *r1);
        println!("r2 is: {}", *r2);
    }
}

pub fn do_unsafe_fn() {
    // 使用不安全块调用不安全函数
    // unsafe fn ..(){} unsafe { .. }
    // 将不安全代码封装进安全函数是一个常见的抽象

    let mut v = vec![1, 2, 3, 4, 5, 6];
    let r = &mut v[..];
    let (a, b) = r.split_at_mut(3);
    assert_eq!(a, &mut [1, 2, 3]);
    assert_eq!(b, &mut [4, 5, 6]);
}

pub fn do_c_extern_fn() {}

pub fn do_modify_static_fn() {}

pub fn do_unsafe_trait() {}

pub fn do_access_union() {
    // 联合体在一个实例中同时只能使用一个声明的字段。联合体主要用于和C代码中的联合体交互
    // 访问联合体的字段是不安全的。因为Rust无法保证当前存储在联合体实例中数据的类型
}
