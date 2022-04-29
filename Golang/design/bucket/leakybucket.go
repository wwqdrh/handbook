package bucket

import (
	"context"
	"fmt"
)

// 每个请求来了，把需要执行的业务逻辑封装成Task，放入木桶，等待worker取出执行

type (
	// 漏桶
	LeakyBucket struct {
		BucketSize int       // 木桶的大小
		NumWorker  int       // 同时从木桶中获取任务执行的worker数量
		bucket     chan Task // 存方任务的木桶
	}

	Task struct {
		handler func() Result // worker从木桶中取出请求对象后要执行的业务逻辑函数
		resChan chan Result   // 等待worker执行并返回结果的channel
	}

	Result struct {
		Data interface{}
	}
)

func NewLeakyBucket(bucketSize int, numWorker int) *LeakyBucket {
	return &LeakyBucket{
		BucketSize: bucketSize,
		NumWorker:  numWorker,
		bucket:     make(chan Task, bucketSize),
	}
}

// 将任务放到漏桶中
func (b *LeakyBucket) Put(handler func() Result) (Result, bool) {
	task := Task{
		handler: handler,
		resChan: make(chan Result),
	}
	// 如果木桶已经满了，返回false
	select {
	case b.bucket <- task:
	default:
		fmt.Println("request is refused")
		return Result{}, false
	}
	// 等待worker执行
	res := <-task.resChan
	return res, true
}

// 启动执行任务
func (b *LeakyBucket) Start(ctx context.Context) {
	// 开启worker从木桶拉取任务执行
	for i := 0; i < b.NumWorker; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("协程任务退出")
					return
				case task := <-b.bucket:
					result := task.handler()
					task.resChan <- result
				}
			}
		}()
	}
}
