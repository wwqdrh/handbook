package gof

import "fmt"

// 桥接模式，多种属性，互相组合，避免子类爆炸

type ICoffee interface {
	OrderCoffee()
}

type ICoffeeAddtion interface {
	AddSomething()
}

type LargeCoffee struct {
	ICoffeeAddtion
}

type MediumCoffee struct {
	ICoffeeAddtion
}

type SmallCoffee struct {
	ICoffeeAddtion
}

// OrderCoffee 订购大杯咖啡
func (lc LargeCoffee) OrderCoffee() {
	fmt.Println("订购了大杯咖啡")
	lc.AddSomething()
}

// OrderCoffee 订购中杯咖啡
func (mc MediumCoffee) OrderCoffee() {
	fmt.Println("订购了中杯咖啡")
	mc.AddSomething()
}

// OrderCoffee 订购小杯咖啡
func (sc SmallCoffee) OrderCoffee() {
	fmt.Println("订购了小杯咖啡")
	sc.AddSomething()
}

type CoffeeCupType uint8

const (
	CoffeeCupTypeLarge  = iota
	CoffeeCupTypeMedium = iota
	CoffeeCupTypeSmall  = iota
)

var CoffeeFuncMap = map[CoffeeCupType]func(coffeeAddtion ICoffeeAddtion) ICoffee{
	CoffeeCupTypeLarge:  NewLargeCoffee,
	CoffeeCupTypeMedium: NewMediumCoffee,
	CoffeeCupTypeSmall:  NewSmallCoffee,
}

// NewCoffee 创建咖啡接口对象的简单工厂，根据咖啡容量类型，获取创建接口对象的func
func NewCoffee(cupType CoffeeCupType, coffeeAddtion ICoffeeAddtion) ICoffee {
	if handler, ok := CoffeeFuncMap[cupType]; ok {
		return handler(coffeeAddtion)
	}
	return nil
}

// NewLargeCoffee 创建大杯咖啡对象
func NewLargeCoffee(coffeeAddtion ICoffeeAddtion) ICoffee {
	return &LargeCoffee{coffeeAddtion}
}

// NewMediumCoffee 创建中杯咖啡对象
func NewMediumCoffee(coffeeAddtion ICoffeeAddtion) ICoffee {
	return &MediumCoffee{coffeeAddtion}
}

// NewSmallCoffee 创建小杯咖啡对象
func NewSmallCoffee(coffeeAddtion ICoffeeAddtion) ICoffee {
	return &SmallCoffee{coffeeAddtion}
}

// Milk 加奶，实现ICoffeeAddtion接口
type Milk struct{}

// Sugar 加糖，实现ICoffeeAddtion接口
type Sugar struct{}

// AddSomething Milk实现加奶
func (milk Milk) AddSomething() {
	fmt.Println("加奶")
}

//AddSomething Sugar实现加糖
func (sugar Sugar) AddSomething() {
	fmt.Println("加糖")
}

// CoffeeAddtionType 咖啡额外添加类型
type CoffeeAddtionType uint8

const (
	// CoffeeAddtionTypeMilk 咖啡额外添加牛奶
	CoffeeAddtionTypeMilk = iota
	// CoffeeAddtionTypeSugar 咖啡额外添加糖
	CoffeeAddtionTypeSugar = iota
)

// CoffeeAddtionFuncMap 全局可导出变量，咖啡额外添加类型与创建咖啡额外添加对象的map，用于减小圈复杂度
var CoffeeAddtionFuncMap = map[CoffeeAddtionType]func() ICoffeeAddtion{
	CoffeeAddtionTypeMilk:  NewCoffeeAddtionMilk,
	CoffeeAddtionTypeSugar: NewCoffeeAddtionSugar,
}

// NewCoffeeAddtion 创建咖啡额外添加接口对象的简单工厂，根据咖啡额外添加类型，获取创建接口对象的func
func NewCoffeeAddtion(addtionType CoffeeAddtionType) ICoffeeAddtion {
	if handler, ok := CoffeeAddtionFuncMap[addtionType]; ok {
		return handler()
	}
	return nil
}

// NewCoffeeAddtionMilk 创建咖啡额外加奶
func NewCoffeeAddtionMilk() ICoffeeAddtion {
	return &Milk{}
}

// NewCoffeeAddtionMilk 创建咖啡额外加糖
func NewCoffeeAddtionSugar() ICoffeeAddtion {
	return &Sugar{}
}
