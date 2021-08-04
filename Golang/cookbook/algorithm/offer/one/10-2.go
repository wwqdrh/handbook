package one

/**
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1

input: 2
output: 2

input: 0
output: 1
*/

func Hand10_2(n int) int {
	if n <= 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, (a+b)%int(1e9+7)
	}
	return b % int(1e9+7)
}
