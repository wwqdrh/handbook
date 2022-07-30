package code

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIpLimit(t *testing.T) {
	limit := NewRateLimit()

	ips := []string{}
	for i := 0; i < 100; i++ {
		ips = append(ips, fmt.Sprintf("1.0.0.%d", i))
	}

	var cnt int64 = 0
	wait := sync.WaitGroup{}
	wait.Add(100)
	for _, item := range ips {
		go func(ip string) {
			defer wait.Done()
			for i := 0; i < 1000; i++ {
				if limit.Do(ip) {
					atomic.AddInt64(&cnt, 1)
				}
			}
		}(item)
	}
	wait.Wait()
	require.Equal(t, int64(100), cnt)
}
