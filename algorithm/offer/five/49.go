package five

import "math"

/**
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

input: 10
output: 12
*/

func min(nums ...int) int {
	minVal := math.MaxInt64
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func Hand49(n int) int {
	idx2, idx3, idx5 := 0, 0, 0
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[idx2]*2, dp[idx3]*3, dp[idx5]*5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			idx2++
		}
		if dp[i] == n3 {
			idx3++
		}
		if dp[i] == n5 {
			idx5++
		}
	}
	return dp[n-1]
}
