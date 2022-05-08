package concurrent

import (
	"fmt"
	"sync"
)

// 使用两个goroutine交替打印序列
// 一个打印数字另外一个答应字母
// 使用channel控制打印进度，数字打印完通知字母打印，字母打印通知数字打印，以此类推
func MultiPrint() {
	letter, number := make(chan bool), make(chan bool)
	wait := &sync.WaitGroup{}
	wait.Add(2)

	go func() {
		i := 1
		for range number {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			letter <- true
		}
		wait.Done()
	}()
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for range letter {
			if i >= 26 {
				wait.Done()
				break
			}
			fmt.Print(str[i : i+1])
			i++

			if i >= 26 {
				wait.Done()
				break
			}
			fmt.Print(str[i : i+1])
			i++
			number <- true
		}
		close(number)
		close(letter)
	}()

	number <- true
	wait.Wait()
}
