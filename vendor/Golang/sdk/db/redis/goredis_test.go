package redis

import "testing"

func TestRedisSet(t *testing.T) {
	StringSet("testkey", "testValue")
}
