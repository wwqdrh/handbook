package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoPool(t *testing.T) {
	wait := sync.WaitGroup{}

	p := NewTaskPool(5)
	for i := 0; i < 10; i++ {
		wait.Add(1)
		p.NewTask(func() {
			defer wait.Done()
			time.Sleep(2 * time.Second)
			fmt.Println(time.Now())
		})
	}

	wait.Wait()
}
