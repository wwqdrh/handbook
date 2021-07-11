package redis

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

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
