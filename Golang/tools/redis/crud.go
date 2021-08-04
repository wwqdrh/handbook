package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

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
