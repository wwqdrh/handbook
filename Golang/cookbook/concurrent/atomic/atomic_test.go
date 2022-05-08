package atomic

import (
	"fmt"
	"sync"
	"testing"
)

func TestSimpleAtomic(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer func() { wait.Done() }()
			BasicWrite()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			defer func() { wait.Done() }()
			fmt.Println(BasicRead())
		}()
	}
	wait.Wait()
	fmt.Println("执行完成", BasicRead())
}
