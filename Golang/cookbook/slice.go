package cookbook

import (
	"fmt"
)

func doArray(a [4]int) {
	a[3] = 10
	fmt.Println(a)
}

func doSlice(s []int) {
	if len(s) > 0 {
		s[0] = 10
	}
	fmt.Println(s)
}

func resSlice(s []int) []int {
	// s = append(s, 5, 6) //append增加切片元素时 相当于创建了一个新的变量
	// fmt.Println(s)
	// return append(s, 20) //返回的是一个新的变量
	s = append(s, 5)
	// fmt.Printf("%v %p %p %v\n", s, s, &s, unsafe.Pointer(&s[0])) // [2 5] 0xc0000dc028 0xc0000c0090 0xc0000dc028
	fmt.Printf("%v %p %p\n", s, s, &s)
	return s
}

func resSlice1(s *[]int) []int {
	*s = append(*s, 5, 6) //这样才能通过append修改元数据的值
	return append(*s, 20)
}

func DoSlice(n int) {
	var a [4]int
	// var a = make([]int, n)
	a[0] = 1
	a[1] = 2
	// fmt.Println(a)

	var b []int
	// // b[0] = 1
	// fmt.Println(b)
	// b = a[:2:4]
	b = a[1:2:4]
	// fmt.Println(cap(b))
	// // c:=a[1:3]
	// // fmt.Println(c) // b c 的底层数组是否同一个
	// // doArray(a)
	// // fmt.Println(a)
	// doSlice(b)
	// // fmt.Println(b)
	fmt.Printf("%v %p %p\n", b, b, &b)
	resSlice(b)
	fmt.Printf("%v %p %p\n", b, b, &b)

	// resSlice1(&b)
	// fmt.Println(b)

}

func DoSlice2() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	fmt.Println(s1, cap(s1)) // 2 3 4 cap:8 如果不指定就跟整个slice长度一样
	s2 := s1[2:6:7]
	fmt.Println(s2, cap(s2)) // 4 5 6 7 cap: 5

	fmt.Println("===")

	s2 = append(s2, 100)
	s2 = append(s2, 200)
	fmt.Println(s2, cap(s2))       // [4 5 6 7 100 200] 10
	fmt.Println(slice, cap(slice)) // [0 1 2 3 4 5 6 7 100 9] 10
	fmt.Println(s1, cap(s1))       // [2 3 4] 8

	fmt.Println("===")

	s1[2] = 20
	fmt.Println(s1, cap(s1))       // 2 3 20 cap:8
	fmt.Println(slice, cap(slice)) // 0, 1, 2, 3, 20, 5, 6, 7, 100, 9 cap: 8

}

func dot(nums []int, a int) []int {
	nums = append(nums, a)
	return nums
}

func DoSlice3(n int) {
	slice1 := []int{}
	for i := 0; i < n; i++ {
		slice1 = append(slice1, i)
	}

	slice2 := append(make([]int, 0, 4), 1, 2, 3)
	fmt.Println(cap(slice2))

	fmt.Printf("slice1: %p %p\n", slice1, &slice1)
	fmt.Printf("slice2: %p %p\n", slice2, &slice2)

	fmt.Println("===")

	fmt.Printf("slice1: %p %p\n", slice1, &slice1)
	slice1 = append(slice1, 100)
	fmt.Printf("slice2: %p %p\n", slice1, &slice1)

	fmt.Printf("slice2: %p %p\n", slice2, &slice2)
	slice2 = append(slice2, 100)
	fmt.Printf("slice2: %p %p\n", slice2, &slice2)
}
