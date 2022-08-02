package concurrent

import (
	"context"
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

func timeout(duration time.Duration, f func(context.Context, chan bool)) {
	// 不用time.After是因为避免没有超时，导致这个timer没有到时间自动退出时协程泄漏状态
	timer := time.NewTimer(duration)
	ctx, cancel := context.WithTimeout(context.TODO(), duration)

	defer func() {
		cancel()
		timer.Stop()
	}()

	done := make(chan bool)
	// 执行的goroutine为发送方
	// 这里没有主动关闭 对于一个 channel，如果最终没有任何 goroutine 引用它，不管 channel 有没有被关闭，最终都会被 gc 回收
	go f(ctx, done)
	select {
	case <-done:
		fmt.Println("done")
	case <-timer.C:
		fmt.Println("timeout")
	}
}

// 这个实现不对，这样还是会导致协程函数中所有代码都全部执行了
func doGoodthing(ctx context.Context, done chan bool) {
	// 模拟长事务 分段执行
	time.Sleep(time.Second)
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Println("执行完了第一段")

	}

	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Println("执行完了第二段")

	}

	time.Sleep(2 * time.Second)
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Println("执行完了第三段")

	}
	time.Sleep(1 * time.Second)
	select {
	case done <- true:
	default: // 使用default防止timeout超时已经退出，这个通道没有接收者了，如果不使用default会一直阻塞
		fmt.Println("全部执行完成")
	}
}

// time waitgroup

func TimeWaitSimple() {
	wg := sync.WaitGroup{}
	c := make(chan struct{}) // close信号
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			select {
			case <-close:
				return
			default: // 模拟业务调用，select外部可嵌套for循环，将业务代码拆成小部分在这default执行
				fmt.Println(num)
			}
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		// 创建的协程捕获到closech关闭，当超时了就主动退出了
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitGroupTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan bool, 1) // 这个通道容量为1就够了，因为有一个会被读出去，不会造成协程阻塞
	done := make(chan struct{})

	// todo 用AfterFunc有内存泄漏风险, 未超时这段时间会一直占着资源
	// go time.AfterFunc(timeout, func() {
	// 	ch <- true
	// })
	go func() {
		timer := time.NewTimer(timeout)
		defer timer.Stop()

		select {
		case <-done:
			return
		case <-timer.C:
			close(done)
			ch <- true
		}
	}()

	go func() {
		wg.Wait()
		select {
		case <-done:
			return
		default:
			// 还未超时
			close(done)
			ch <- false
		}
	}()

	return <-ch
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

// 交替打印数字和字母
func crossPrint() {
	letterCh, numberCh := make(chan struct{}), make(chan struct{})

	wait := sync.WaitGroup{}
	wait.Add(2)

	go func() {
		defer wait.Done()
		i := 1
		for range numberCh {
			fmt.Print(i)
			letterCh <- struct{}{}
			i++
		}
	}()

	go func() {
		defer wait.Done()
		ch := 'A'
		for range letterCh {
			fmt.Print(string(ch))
			if ch == 'Z' {
				close(numberCh)
				break
			}
			numberCh <- struct{}{}
			ch++
		}
	}()

	numberCh <- struct{}{}
	wait.Wait()
}
