package seqgen

import (
	"math"
	"reflect"
)

////////////////////
// 一个handler需要提供
// 1、嵌套处理的能力，也就是本身不具备规律但是按照规则处理后的序列能够具备
// 2、满足规则后执行的生成下一个元素的功能
////////////////////
type ListHandler func(nums []int) (bool, []int, func(a, b int) int)

func addListHandler(nums []int) (bool, []int, func(a, b int) int) {
	res := make([]int, 0, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		res = append(res, nums[i]-nums[i-1])
	}
	return true, res, func(a, b int) int { return a + b }
}

func mulListHandler(nums []int) (bool, []int, func(a, b int) int) {
	res := make([]int, 0, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == 0 {
			return false, nil, nil
		}
		temp := nums[i] / nums[i-1]
		if temp*nums[i-1] != nums[i] {
			return false, nil, nil
		}
		res = append(res, temp)
	}
	return true, res, func(a, b int) int { return a * b }
}

func sqrtListHandler(nums []int) (bool, []int, func(a, b int) int) {
	res := make([]int, 0, len(nums))
	for _, item := range nums {
		if item < 0 {
			return false, nil, nil
		}
		temp := int(math.Sqrt(float64(item)))
		if temp*temp != item {
			return false, nil, nil
		}
		res = append(res, temp)
	}
	if reflect.DeepEqual(res, nums) {
		return false, nil, nil
	}

	return true, res, func(a, b int) int { return b * b }
}

// 按照自定义的优先级组合，可扩展
var combination = []ListHandler{addListHandler, mulListHandler, sqrtListHandler}

// 主函数
func checkList(nums []int) (int, bool) {
	if len(nums) < 3 {
		return 0, false
	}

	var check func(nums []int) (int, bool)
	check = func(nums []int) (int, bool) {
		// 当只剩下一个元素就表示没有找到规律
		if len(nums) <= 1 {
			return 0, false
		}
		allEqual := true
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				allEqual = false
			}
		}
		if allEqual {
			return nums[0], true
		}

		for _, hand := range combination {
			status, newList, cb := hand(nums)
			if !status {
				continue
			}

			next, ok := check(newList)
			if ok {
				return cb(nums[len(nums)-1], next), true
			}
		}
		return 0, false
	}

	return check(nums)
}
