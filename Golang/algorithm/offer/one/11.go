package one

/**
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

input: [3,4,5,1,2]
output: 1

input: [2,2,2,0,1]
output: 0
*/

func Hand11(numbers []int) int {
	// 一般有序数组需要想到的就是二分法，这里二分法需要找到的元素满足必须小于左边并且小于等于右边
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + ((right - left) >> 1)
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right-- // 可能遇到相等的元素，取左边
		}
	}
	return numbers[left]
}
