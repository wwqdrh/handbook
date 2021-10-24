package main

import (
	"fmt"
)

func main() {
	var array [10]int
	var slice = array[5:6:6]
	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])

	slice2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	println(cap(slice2))
	s1 := slice2[2:5]
	println(cap(s1))
	fmt.Printf("%v %v\n", slice2, s1)
	s1 = append(s1, 10)
	fmt.Printf("%v %v\n", slice2, s1)

	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}
