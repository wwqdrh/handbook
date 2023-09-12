package redis

import (
	"fmt"
	"os"
	"time"

	redis "gopkg.in/redis.v5"
)

// Setup initializes a redis client
func Setup() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0, // use default DB
	})

	// commands returns "PONG", tests if
	// connection alive
	_, err := client.Ping().Result()
	return client, err
}

// Exec performs some redis operations
func Exec() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	c1 := "value"
	// value is an interface, we can store whatever
	// the last argument is the redis expiration
	conn.Set("key", c1, 5*time.Second)

	var result string
	if err := conn.Get("key").Scan(&result); err != nil {
		switch err {
		// this means the key
		// was not found
		case redis.Nil:
			return nil
		default:
			return err
		}
	}

	fmt.Println("result =", result)

	return nil
}

// Sort performs a sort redis operations
func Sort() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	if err := conn.LPush("list", 1).Err(); err != nil {
		return err
	}
	if err := conn.LPush("list", 3).Err(); err != nil {
		return err
	}
	if err := conn.LPush("list", 2).Err(); err != nil {
		return err
	}

	res, err := conn.Sort("list", redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	conn.Del("list")
	return nil
}
