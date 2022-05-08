package concurrent

import (
	"bytes"
	"fmt"
	"sync"
)

// 模拟两个协程的通信 一个写一个读并打印出来
func MockCommunicate() string {
	out := make(chan int)
	buf := bytes.NewBufferString("")
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			// out <- rand.Intn(5)
			out <- 1
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for num := range out {
			buf.WriteString(fmt.Sprint(num))
		}
	}()

	wg.Wait()

	return buf.String()
}
