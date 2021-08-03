package gredis

import (
	"archbook/restapi/pkg/setting"
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisConn *redis.Client

func Setup() error {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password,
		// MaxRetries: setting.RedisSetting,Max,
		PoolSize: setting.RedisSetting.MaxIdle,
		// MinIdleConns: ,
	})
	if err := RedisConn.Ping(context.TODO()).Err(); err != nil {
		return err
	}
	return nil
}
