package algorithm

import (
	"math/rand"
	"testing"
)

func BenchmarkSortArrayByParity(b *testing.B) {
	nums := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums = append(nums, rand.Intn(100000))
	}

	for n := 0; n < b.N; n++ {
		sortArrayByParity(nums)
	}
}

func BenchmarkSortArrayBySort(b *testing.B) {
	nums := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums = append(nums, rand.Intn(100000))
	}

	for n := 0; n < b.N; n++ {
		sortArrayBySort(nums)
	}
}
