package five

import "strconv"

/**
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

input: 12258
output: 5 // "bccfi", "bwfi", "bczi", "mcfi"和"mzi"
*/

func Hand46(num int) int {
	// dp[i]表示前i个元素能够翻译的种类 dp[i] = |-- dp[i-1] + dp[i-2]  如果num[i-1]与num[i-2]能够组成小于26
	//										|
	//										|-- dp[i-1]
	numStr := strconv.FormatInt(int64(num), 10)
	length := len(numStr)

	a, b := 1, 1 // dp简化
	for i := 2; i <= length; i++ {
		if numStr[i-2:i] >= "10" && numStr[i-2:i] <= "25" {
			a, b = b, a+b
		} else {
			a = b
		}
	}
	return b
}
