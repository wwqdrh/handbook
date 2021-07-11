package redis

import (
	"context"
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
