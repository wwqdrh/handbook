package cache

import (
	"fmt"
)

func ExampleLruCache() {
	cache := NewLruCache(5)
	cache.Add(1, 1)
	cache.Add(2, 2)
	cache.Add(3, 3)
	cache.Add(4, 4)
	cache.Add(5, 5)
	cache.Add(6, 6)
	for _, item := range cache.GetAll() {
		fmt.Println(item.Val.(int)) // 不能使用println 必须使用fmt包中的
	}

	// Unordered output: 2
	// 3
	// 4
	// 5
	// 6
}
