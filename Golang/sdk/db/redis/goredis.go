package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

func init() {
	// connection string or redis.Options
	// redis://<user>:<pass>@localhost:6379/<db>
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0, // default DB,
		// TLSConfig: &tls.Config{},
	})
}

func testConn() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis连接失败: ", pong, err)
		return err
	} else {
		fmt.Println("redis连接成功: ", pong)
		return nil
	}
}

////////////////////
// pipeline
////////////////////

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

////////////////////
// pubsub
////////////////////

func TryPub(channel, message string) {
	err := client.Publish(ctx, channel, message).Err()
	if err != nil {
		fmt.Println("发生错误")
	}
}

func TrySub(channel string) string {
	sub := client.Subscribe(ctx, channel)
	defer sub.Close()

	// for {
	// 	msg, err := sub.ReceiveMessage(ctx)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(msg.Channel, msg.Payload)
	// }
	ch := sub.Channel()
	var message string
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		message = msg.Payload
	}
	return message
}

////////////////////
// transaction
////////////////////

func Transaction(key string) error {
	maxRetries := 1000
	txf := func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int() // key不存在的时候值为0
		if err != nil && err != redis.Nil {
			return err
		}
		n++
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})
		return err
	}
	for i := 0; i < maxRetries; i++ {
		err := client.Watch(ctx, txf, key)
		if err == nil {
			return nil
		}
		if err == redis.TxFailedErr {
			continue
		}
		return err
	}
	return errors.New("increment reached maximum number of retries")
}

////////////////////
// crud
////////////////////

// string
func StringGet(key string) string {
	if err := testConn(); err != nil {
		fmt.Println("redis尚未连接")
		return ""
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	val, err := client.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key dose not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}
	return val
}

func StringSet(key, value string) bool {
	if err := testConn(); err != nil {
		fmt.Println("redis尚未连接")
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := client.Set(ctx, key, value, 5*time.Second).Err()
	if err != nil {
		fmt.Printf("set failed, err:%v\n", err)
		return false
	}
	return true
}

func StringDel(key string) bool {
	if err := testConn(); err != nil {
		fmt.Println("redis尚未连接")
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Del(ctx, key).Result()
	if err != nil {
		fmt.Println("redis删除失败")
		return false
	}
	return true
}

// list

// set

// zset

// hash
