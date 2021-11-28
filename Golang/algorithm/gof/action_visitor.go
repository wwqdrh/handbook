package gof

import "fmt"

// 抽象访问者
// 访问者
// 抽象元素类
// 元素类
// 结构容器
// 自从各个语言开始支持匿名函数之后，访问者模式就变得极其简单了，每一种传入匿名方法的操作都可以看做是变相的访问者模式， golang 中的方法也是一种类型的对象，所以可以用它便利的实现访问者模式

// 定义访问者接口
type IVisitor interface {
	Visit() // 访问者的访问方法
}

// 定义元素接口
type IElement interface {
	Accept(visitor IVisitor)
}

type ProductionVisitor struct {
}

func (v ProductionVisitor) Visit() {
	fmt.Println("这是生产环境")
}

type TestingVisitor struct {
}

func (t TestingVisitor) Visit() {
	fmt.Println("这是测试环境")
}

type Element struct{}
type EnvExample struct {
	Element
}

func (el Element) Accept(visitor IVisitor) {
	visitor.Visit()
}

func (e EnvExample) Print(visitor IVisitor) {
	e.Element.Accept(visitor)
}
