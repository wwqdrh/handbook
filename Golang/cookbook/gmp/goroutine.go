package gmp

////////////////////
// 测试栈扩容
////////////////////

//go:noinline
func a() {
	println(`func a`)
	var y [1000]int
	b(y)
}

//go:noinline
func b(x [1000]int) {
	println(`func b`)
	var y [100000]int
	c(y)
}

//go:noinline
func c(x [100000]int) {
	println(`func c`)
}
