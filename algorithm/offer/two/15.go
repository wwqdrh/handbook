package two

/**
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。

input: 11
output: 3
*/

func Hand15(num uint32) int {
	// bytes.OnesCount(int(num))
	res := 0
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
