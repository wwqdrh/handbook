package algorithm_test

import (
	"reflect"
	"testing"
	"wwqdrh/handbook/algorithm"
)

func TestCountSort(t *testing.T) {
	nums := []int{4, 5, 6, 7, 2, 10, 2}
	numsSort := algorithm.CountSort(nums)

	if !reflect.DeepEqual(nums, []int{4, 5, 6, 7, 2, 10, 2}) {
		t.Error("函数对外部元素产生了影响")
	}
	if !reflect.DeepEqual(numsSort, []int{2, 2, 4, 5, 6, 7, 10}) {
		t.Error("CountSort排序失败")
	}
}

func TestBucketSort(t *testing.T) {
	nums := []int{4, 5, 6, 7, 2, 10, 2}
	numsSort := algorithm.BucketSort(nums)

	if !reflect.DeepEqual(nums, []int{4, 5, 6, 7, 2, 10, 2}) {
		t.Error("函数对外部元素产生了影响")
	}
	if !reflect.DeepEqual(numsSort, []int{2, 2, 4, 5, 6, 7, 10}) {
		t.Error("CountSort排序失败")
	}
}

func TestRadixSort(t *testing.T) {
	nums := []int{4, 5, 6, 7, 2, 10, 2, 19, 25, 1, 12, 54}
	numsSort := algorithm.RadixSort(nums)

	if !reflect.DeepEqual(nums, []int{4, 5, 6, 7, 2, 10, 2, 19, 25, 1, 12, 54}) {
		t.Error("函数对外部元素产生了影响")
	}
	if !reflect.DeepEqual(numsSort, []int{1, 2, 2, 4, 5, 6, 7, 10, 12, 19, 25, 54}) {
		t.Error("RadixSort排序失败")
	}

	nums = []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	numsSort = algorithm.RadixSort(nums)

	if !reflect.DeepEqual(nums, []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}) {
		t.Error("函数对外部元素产生了影响")
	}
	if !reflect.DeepEqual(numsSort, []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}) {
		t.Error("RadixSort排序失败")
	}
}
