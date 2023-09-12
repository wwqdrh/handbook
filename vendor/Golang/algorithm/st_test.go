package algorithm

import "testing"

func TestSumSubarrayMins(t *testing.T) {
	if sumSubarrayMins([]int{3, 1, 2, 4}) != 17 {
		t.Error("子数组和失败")
	}
}
