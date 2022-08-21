package main

import (
	"fmt"
)

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func GetValue() int {
	return 1
}

func typeselect() {
	var i interface{} = GetValue()
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}

func hello(num ...int) {
	num[0] = 18
}

type T struct {
	n int
}

func getT() T {
	return T{}
}

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func alwaysFalse() bool {
	return false
}

func main() {
	aa0, aa, aa1, aa2, aa3 := byte(0), byte(1), byte(128), byte(64), byte(129)
	fmt.Println(aa, -aa, aa0, -aa0, aa1, -aa1, aa2, -aa2, aa3, -aa3)
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			fmt.Println(n)
			count++
		}
		if m == -m {
			fmt.Println(m)
			count++
		}
	}
	fmt.Println(count)

	// 有分号和分号结果不一样😹
	switch alwaysFalse(); {
	case true:
		println(true)
	case false:
		println(false)
	}
	funs := test()
	for _, f := range funs {
		f()
	}
	var x int8 = -128
	var y = x / -1
	fmt.Println(x, y)

	var x1 int8 = -128
	var y1 = x1/-1 - 1
	fmt.Println(x1, y1)

	t1 := getT()
	// p := &t1.n // <=> p = &(t.n)
	// *p = 1
	t1.n = 1
	fmt.Println(t1.n)

	s := make([]int, 3, 9)
	fmt.Println(len(s))
	s2 := s[4:8]
	fmt.Println(len(s2), s2)

	fmt.Println(^2)

	a1 := 1
	if a1, b1 := 2, 2; a1 == 2 {
		fmt.Println(a1, b1)
	}
	fmt.Println(a1)
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])

	nums := [5]int{1, 2, 3, 4, 5}
	t := nums[3:4:4]
	fmt.Println(t[0], cap(t))

	typeselect()
	fmt.Println(x, y, z, k, p)

	fmt.Println(0.1+0.2 == 0.3)
	fmt.Println(18466.67 * 100)

	a := new([]int)
	b := *a
	b = append(b, 1)
	fmt.Println(b)
}
