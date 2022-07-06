package compiler

import (
	"fmt"
	"time"
)

func Sort() {
	start := time.Now().UnixNano()
	// var arr []int
	const NUM int = 100000000
	for i := 0; i < NUM; i++ {
		//    arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 这样就不会逃逸了 神奇，可能是与下面的优化成一行了 没有在当前的函数栈中申请空间
		bubbleSort(arr)
	}
	//打印消耗时间
	fmt.Println(time.Now().UnixNano() - start)
}

func Sort2() {
	start := time.Now().UnixNano()
	var arr []int
	const NUM int = 100000000
	for i := 0; i < NUM; i++ {
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		bubbleSort(arr)
	}
	//打印消耗时间
	fmt.Println(time.Now().UnixNano() - start)
}

//排序
func bubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-1-j; k++ {
			if arr[k] < arr[k+1] {
				temp := arr[k]
				arr[k] = arr[k+1]
				arr[k+1] = temp
			}
		}
	}
}

func escapeint() {
	var a int64
	for i := 0; i < 100; i++ {
		a = 10
		escapecall(&a) // 就算是int这种可复制的类型，这样申明也会被逃逸到堆中
	}
}

func escapecall(a *int64) {
	fmt.Println(*a)
}
