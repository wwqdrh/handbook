package one

/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

input: [2, 3, 1, 0, 2, 5, 3]
output: [2, 3]
*/

func Hand3(nums []int) int {
	visit := make(map[int]bool)
	for _, num := range nums {
		if !visit[num] {
			visit[num] = true
		} else {
			return num
		}
	}
	return -1
}
