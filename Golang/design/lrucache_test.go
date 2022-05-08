package design

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
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

func ExampleConsumerQueue() {
	queue := NewConsumerQueue()
	res := []int{}
	mutex := new(sync.Mutex)

	wg := sync.WaitGroup{}
	wg.Add(3)

	// 生产者
	go func() {
		for i := 1; i <= 5; i++ {
			queue.Producer(i)
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), (1 * time.Second))
		defer cancel()
		if err := queue.Close(timeoutCtx); err != nil {
			fmt.Printf("err: %v \n", err)
		}
		wg.Done()
	}()

	// 消费者
	go func() {
		for {
			if i, ok := queue.Consumer(); ok {
				mutex.Lock()
				res = append(res, i.(int))
				mutex.Unlock()
			} else {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		for {
			if i, ok := queue.Consumer(); ok {
				mutex.Lock()
				res = append(res, i.(int))
				mutex.Unlock()
			} else {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()

	for _, item := range res {
		fmt.Println(item)
	}

	// Unordered output: 1
	// 2
	// 3
	// 4
	// 5
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go OnceInstance()
	}
}
