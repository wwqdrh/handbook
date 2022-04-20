package design

import (
	"context"
	"errors"
	"sync/atomic"
	"time"
)

// 需求分析
// 多个协程写 多个协程等待读
// 消息不能重复消费

type ConsumerQueue struct {
	count   int64
	channel chan interface{}
}

func NewConsumerQueue() *ConsumerQueue {
	return &ConsumerQueue{
		count:   0,
		channel: make(chan interface{}, 10),
	}
}

func (q *ConsumerQueue) Producer(v interface{}) {
	q.channel <- v
	atomic.AddInt64(&q.count, 1)
}

func (q *ConsumerQueue) Consumer() (interface{}, bool) {
	val, ok := <-q.channel
	atomic.AddInt64(&q.count, -1)
	return val, ok
}

func (q *ConsumerQueue) Close(ctx context.Context) error {
	// 如果直接唤醒 读的直接唤醒写入零值
	// 写的直接panic

	for {
		select {
		case <-ctx.Done():
			return errors.New("超时")
		default:
			if q.count == 0 {
				close(q.channel)
				return nil
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}
