package two

import "math"

/**
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

input: 1
output: [1,2,3,4,5,6,7,8,9]
*/

func Hand17(n int) []int {
	res := make([]int, 0)
	maxNum := int(math.Pow(float64(10), float64(n))) - 1
	for i := 1; i <= maxNum; i++ {
		res = append(res, i)
	}
	return res
}
