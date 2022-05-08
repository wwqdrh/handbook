package math

import "fmt"

func IsInt(bits uint32, bias int) {
	// 计算指数
	exponent := int(bits>>23) - bias - 23
	// 通过位运算方式计算小数部分的值
	coefficient := (bits & ((1 << 23) - 1)) | (1 << 23)
	// 计算intTest 用于判断指数能否弥补小数部分
	intTest := (coefficient & (1<<uint32(-exponent) - 1))
	// fmt.Printf("\nExponent:%d Coefficient: %d IntTest: %d\n",
	// 	exponent, coefficient, intTest,
	// )
	if exponent < -23 {
		fmt.Println("NOT INTEGER")
		return
	}
	if exponent < 0 && intTest != 0 {
		fmt.Println("NOT INTEGER")
		return
	}
	fmt.Println("INTEGER")
}
