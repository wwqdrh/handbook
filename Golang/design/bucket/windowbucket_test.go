package bucket

import (
	"fmt"
	"testing"
	"time"
)

func TestWindowBucket(t *testing.T) {
	limiter := NewSliding(100*time.Millisecond, time.Second, 10)
	for i := 0; i < 5; i++ {
		fmt.Println(limiter.IsLimited())
	}
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 5; i++ {
		fmt.Println(limiter.IsLimited())
	}
	fmt.Println(limiter.IsLimited())
	for _, v := range limiter.windows[limiter.getUidOrIp()] {
		fmt.Println(v.timestamp, v.count)
	}

	fmt.Println("a thousand years later...")
	time.Sleep(time.Second)
	for i := 0; i < 7; i++ {
		fmt.Println(limiter.IsLimited())
	}
	for _, v := range limiter.windows[limiter.getUidOrIp()] {
		fmt.Println(v.timestamp, v.count)
	}

}
