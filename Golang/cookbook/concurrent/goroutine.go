package concurrent

import (
	"fmt"
	"log"
	"sync"
	"time"
)

////////////////////
// 超时退出协程
////////////////////

// 无法强制kill一个协程
// 1、尽量使用非阻塞I/O
// 2、业务逻辑总是考虑退出机制
// 3、任务分段执行，超时后即使退出

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default: // 使用default防止timeout超时已经退出，这个通道没有接收者了，如果不使用default会一直阻塞
		return
	}
}

////////////////////
// 确保通道只被关闭一次
////////////////////
// 情形一：M个接收者和一个发送者，发送者通过关闭用来传输数据的通道来传递发送结束信号。
// 情形二：一个接收者和N个发送者，此唯一接收者通过关闭一个额外的信号通道来通知发送者不要再发送数据了。
// 情形三：M个接收者和N个发送者，它们中的任何协程都可以让一个中间调解协程帮忙发出停止数据传送的信号。

type MyChannel struct {
	C    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{
		C: make(chan int),
	}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

////////////////////
// 限制协程并发数量
////////////////////
func LimitGoroutineNumber() {
	var wg sync.WaitGroup // 用于等待最后三个
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{} // 当写了三次就不能再写了，除非下文有程序执行完了能够继续写入
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}
