package bucket

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	context, cancel := context.WithCancel(context.Background())
	bucket := NewLeakyBucket(3, 1)
	bucket.Start(context)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		if res, ok := bucket.Put(func() Result {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("执行任务中...")
			return Result{"done"}
		}); !ok {
			fmt.Println("放入失败")
		} else {
			fmt.Println("放入成功, 结果为", res)
		}
		wg.Done()
	}
	wg.Wait()
	cancel()
}
