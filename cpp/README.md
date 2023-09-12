# usage


```bash
wget https://xmake.io/shget.text -O - | bash

xmake project -k compile_commands

xmake build test_os

xmake run test_os
```

# 语法

c++17, 才有std::filesystem

1、包含基础语法
2、各个版本之间的新增语法

```bash
g++ [file] -o [输出]

// 执行输出文件
```

- 数据类型：long long，char16_t，char32_t, char8_t
- 内联和嵌套命名空间
- auto占位符
- decltype
- 函数返回类型后置
- 右值引用
- lambda表达式
- 非静态数据成员默认和初始化
- 列表初始化
- 默认和删除函数
- 非受限联合类型
- 委托构造函数
- 继承构造函数
- 强枚举类型
- 扩展的聚合类型
- override和final说明符
- 基于范围的for循环
- 支持初始化语句的if和switch
- static_assert声明
- 结构化绑定
- noexcept关键字
- 类型别名和别名模版
- 指针字面量nullptr
- 三向比较
- 线程局部存储
- 扩展的inline说明符
- 常量表达式
- 确定的表达式求值顺序
- 字面量优化
- alignas和alignof
- 属性说明符和标准属性
- 新增预处理器和宏
- 协程
- 可变参数模版
- typename优化
- 模版参数优化
- 类模版的模版实参推导
- 用户自定义推导指引
- SFINAE
- 概念和约束
- 模版特性的其他优化

