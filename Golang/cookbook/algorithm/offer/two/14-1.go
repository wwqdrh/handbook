package two

import "math"

/**
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

input: 2
output: 1

input: 10
output: 36
*/

func Hand14_1(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	parts, another := n/3, n%3

	var result float64
	switch another {
	case 2:
		result = math.Pow(3, float64(parts))
		result *= 2
	case 1: // 3*1转化为2*2
		result = math.Pow(3, float64(parts-1))
		result *= 4
	default:
		result = math.Pow(3, float64(parts))
	}
	return int(result)
}
