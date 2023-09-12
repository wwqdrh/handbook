package redis

import "testing"

func TestRedis5(t *testing.T) {
	if err := Exec(); err != nil {
		panic(err)
	}

	if err := Sort(); err != nil {
		panic(err)
	}
}
