package concurrent

import (
	"bytes"
	"fmt"
	"sync"
)

// 使用两个goroutine交替打印序列
// 一个打印数字另外一个答应字母
// 使用channel控制打印进度，数字打印完通知字母打印，字母打印通知数字打印，以此类推
func MultiPrint() string {
	outBuf := bytes.NewBufferString("")

	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				outBuf.WriteString(fmt.Sprint(i))
				i++
				outBuf.WriteString(fmt.Sprint(i))
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= 26 {
					wait.Done()
					return
				}
				outBuf.WriteString(str[i : i+1])
				i++

				if i >= 26 {
					wait.Done()
					return
				}
				outBuf.WriteString(str[i : i+1])
				i++
				number <- true
			}
		}
	}(&wait)

	number <- true
	wait.Wait()
	return outBuf.String()
}
