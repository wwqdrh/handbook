package one

/**
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。
斐波那契数列的定义如下：
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

input: 2
output: 1

input: 5
output: 5
*/

func Hand10_1(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
