package two

/**
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。
请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

input: 10
output: 36
*/

func Hand14_2(n int) int {
	// 大数越界下的求余问题， 快速幂求余
	if n <= 3 {
		return n - 1
	}
	b := n % 3
	rem, x, p := int64(1), int64(3), int64(1e9+7)
	for a := n/3 - 1; a > 0; a >>= 1 {
		if (a & 1) == 1 {
			rem = (x * rem) % p
		}
		x = (x * x) % p
	}
	if b == 0 {
		return (int)(rem * 3 % p)
	} else if b == 1 {
		return (int)(rem * 4 % p)
	}
	return (int)(rem * 6 % p)
}
