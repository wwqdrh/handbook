package concurrent

import (
	"runtime"
	"testing"
	"time"
)

func TestGroutinTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeout(doGoodthing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}
