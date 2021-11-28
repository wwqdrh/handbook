package gof

// 结构型-工厂
// 抽象工厂
// 生成器
// 原型
// 单例

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

type PlusOperator struct {
	*OperatorBase
}

type MinusOperator struct {
	*OperatorBase
}

type PlusOperatorFactory struct{}

type MinusOperatorFactory struct{}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

func FactoryMode(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
