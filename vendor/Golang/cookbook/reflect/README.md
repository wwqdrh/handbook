反射与接口强相关

# 原理

为将传递进来的接口变量转换为底层的实际空接口 emptyInterface，并获取空接口的类型值。

# reflect.Type

1、reflect.TypeOf获取
2、reflect.Value.Type获取

- Name: 类型的名字
- Align
- Method

# reflect.Value

由reflect.ValueOf获取、

- Elem: 如果当前的值是指针，通过Elem获取指针指向的值
- Interface
- Int
- String
- ...

``` go
var z = 123
var y = &z
var x interface{} = y

v := reflect.ValueOf(&x)
vx := v.Elem()
fmt.Println(vx.Kind()) // interface{}
vy := vx.Elem()
fmt.Println(vy.Kind()) // ptr
vz := vy.Elem()
fmt.Println(vz.Kind()) // int
```

# reflect.StructField

由`reflect.Type.Field获取`

- Tag: 字段tag