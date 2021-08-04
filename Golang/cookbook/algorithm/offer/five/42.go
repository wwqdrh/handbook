package five

import "math"

/**
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。
*/

func Hand42(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, res := 0, math.MinInt64
	for _, num := range nums {
		pre = max(pre+num, num)
		res = max(res, pre)
	}
	return res
}
