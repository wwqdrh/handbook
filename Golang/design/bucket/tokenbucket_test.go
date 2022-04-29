package bucket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTokenbucket(t *testing.T) {
	tokenBucket := NewTokenBucket(5, 100*time.Millisecond)
	for i := 0; i < 6; i++ {
		fmt.Println(tokenBucket.IsLimited())
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(tokenBucket.IsLimited())
}

func TestTokenConcurrentPut(t *testing.T) {
	tokenBucket := NewTokenBucket(5, 100*time.Millisecond)
	wait := sync.WaitGroup{}
	wait.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			tokenBucket.Put("127.0.0.1")
			wait.Done()
		}()
	}
	wait.Wait()
}
