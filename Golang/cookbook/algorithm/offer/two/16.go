package two

/**
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。

intput: x=2.0 n=10
output: 1024.0
*/

func Hand16(x float64, n int) float64 {
	res := float64(1)
	if n < 0 {
		x, n = float64(1/x), -n
	}
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1 // 快速幂
	}
	return res
}
