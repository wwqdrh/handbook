package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Pipeline 主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。
// 这些命令不能保证在事务中执行。
// 这样做的好处是节省了每个命令的网络往返时间（RTT）。

func PipelineIncr() {
	if err := testConn(); err != nil {
		fmt.Println("redis尚未连接")
		return
	}
	ctx := context.Background()

	// 1、
	var incr *redis.IntCmd
	_, err := client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipelined_counter")
		pipe.Expire(ctx, "pipelined_counter", time.Second*5)
		return nil
	})
	if err != nil {
		fmt.Printf("pipeline_counter err: %v", err)
	} else {
		fmt.Println(incr.Val())
	}

	// 2、
	pipe := client.Pipeline()
	incr = pipe.Incr(ctx, "pipeline_counter2")
	pipe.Expire(ctx, "pipeline_counter2", time.Second*5)
	_, err = pipe.Exec(ctx)
	if err != nil {
		fmt.Printf("pipeline_counter2 err: %v", err)
	} else {
		fmt.Println(incr.Val())
	}
}
