package gmp

import (
	"fmt"
	"testing"
)

func TestGoroutineStack(t *testing.T) {
	// wired! 这里打印出来的地址是一样的
	// 破案了，fmt.Println内部应该有协程共享了变量，这个变量是被分配到了堆上
	// 测试
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	println(&x)
	// }()
	// wg.Wait()
	var x [10]int
	println(&x)
	// fmt.Println(&x)
	fmt.Printf("%p \n", &x)
	a()
	println(&x)
	fmt.Printf("%p \n", &x)

}

func TestGoroutineStack2(t *testing.T) {
	var x [10]int
	println(&x)
	a()
	println(&x)
}
