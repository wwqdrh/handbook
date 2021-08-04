package five

import "strconv"

/**
数字以0123456789101112131415…的格式序列化到一个字符序列中。
在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。
*/

func Hand44(n int) int {
	// 先判断在哪一个层级，第一层级有9个，第二层级有180，第三层级3 * 100 * 9
	// 在哪一个数字，长度为3的在第二层级那就是第二个数
	// 确定在哪一位, 3在2层，则取第二个数的第一位
	digit, start, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = digit * start * 9
	}
	num, dig := (n-1)/digit, (n-1)%digit

	res, _ := strconv.ParseInt(string(strconv.FormatInt(int64(start+num), 10)[dig]), 10, 0)
	return int(res)
}
