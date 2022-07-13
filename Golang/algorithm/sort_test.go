package algorithm

import (
	"fmt"
	"math/rand"
	"sort"
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

func TestMutilFieldSort(t *testing.T) {
	type record struct {
		Name  string
		Value int
	}

	records := []record{
		{"user1", 1},
		{"user1", 4},
		{"user3", 6},
		{"user5", 8},
		{"user2", 9},
		{"user4", 3},
		{"user5", 2},
		{"user6", 4},
		{"user7", 6},
	}

	sort.Slice(records, func(i, j int) bool {
		a, b := records[i], records[j]
		if a.Name < b.Name {
			return true
		} else if a.Name > b.Name {
			return false
		} else {
			return a.Value < b.Value
		}
	})

	fmt.Println(records)
}
