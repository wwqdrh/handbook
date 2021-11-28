package gof

import "fmt"

// 适配 组合 桥接 装饰 外观 享元 单例

// 源结构体
type Adaptee interface {
	OutputPower()
}

// 目标结构体
type Target interface {
	CovertTo5V()
}

// 具体结构体
type Volts220 struct{}

func (v Volts220) OutputPower() {
	fmt.Println("电源输出了220V电压")
}

type Adapter struct {
	Adaptee
}

func (a Adapter) CovertTo5V() {
	// 自定义的转换逻辑
	a.OutputPower()
	fmt.Println("通过手机电源适配器，转成了5V电压，可供手机充电")
}
