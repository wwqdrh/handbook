package two

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

input: [1,2,3,4]
output: [1,3,2,4]
*/

func Hand21(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]&1 == 0 && nums[right]&1 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left]&1 == 1 {
			left++
		} else if nums[right]&1 == 0 {
			right--
		}
	}
	return nums
}
