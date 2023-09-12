package cookbook

import "fmt"

type X struct{}

func (x *X) tf() {
	println(x)
}

func StructSimple() {
	var a *X
	a.tf()

	x := X{}
	x.tf()
}

type N int

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%p,%v\n", n, *n)
}

func StructCall() {

	var a N = 25

	p := &a
	// p1 := &p
	p.value()
	p.pointer()

	// p1.value() // 不能使用多级指针调用方法
	// p1.pointer() // 不能使用多级指针调用方法
}

type N1 int

func (n N1) test() {
	fmt.Println(n)
}

func plainFunc() {
	var n N1 = 10
	fmt.Println(n)

	n++
	f1 := N1.test
	f1(n)

	n++
	f2 := (*N1).test
	f2(&n)
}
