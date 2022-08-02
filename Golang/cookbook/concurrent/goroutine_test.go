package concurrent

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGroutinTimeout(t *testing.T) {
	wait := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			timeout(5*time.Second, doGoodthing)
		}()
	}
	wait.Wait()
	t.Log(runtime.NumGoroutine())
}

func TestCrossPrint(t *testing.T) {
	crossPrint()
}
