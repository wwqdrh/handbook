package concurrent

// 常规的通过通道发送和接收数据是阻塞的。
// 然而，我们可以使用带一个 `default` 子句的 `select`
// 来实现 _非阻塞_ 的发送、接收，甚至是非阻塞的多路 `select`。

import (
	"fmt"
	"time"
)

func ChannelExample() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这是一个非阻塞接收的例子。
	// 如果在 `messages` 中存在，然后 `select` 将这个值带入 `<-messages` `case` 中。
	// 否则，就直接到 `default` 分支中。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 `default` 前使用多个 `case` 子句来实现一个多路的非阻塞的选择器。
	// 这里我们试图在 `messages` 和 `signals` 上同时使用非阻塞的接收操作。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func ChannelExample2() {

	// 我们将遍历在 `queue` 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// `range` 迭代从 `queue` 中得到每个值。
	// 因为我们在前面 `close` 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。
	for elem := range queue {
		fmt.Println(elem)
	}
}

func ChannelBuffer() {

	// 这里我们 `make` 了一个字符串通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// 由于此通道是有缓冲的，
	// 因此我们可以将这些值发送到通道中，而无需并发的接收。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// `ping` 函数定义了一个只能发送数据的（只写）通道。
// 尝试从这个通道接收数据会是一个编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` 函数接收两个通道，`pings` 仅用于接收数据（只读），`pongs` 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirection() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// _关闭_ 一个通道意味着不能再向这个通道发送值了。
// 该特性可以向通道的接收方传达工作已经完成的信息。

// 在这个例子中，我们将使用一个 `jobs` 通道，将工作内容，
// 从 `main()` 协程传递到一个工作协程中。
// 当我们没有更多的任务传递给工作协程时，我们将 `close` 这个 `jobs` 通道。
func CloseChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是工作协程。使用 `j, more := <- jobs` 循环的从 `jobs` 接收数据。
	// 根据接收的第二个值，如果 `jobs` 已经关闭了，
	// 并且通道中所有的值都已经接收完毕，那么 `more` 的值将是 `false`。
	// 当我们完成所有的任务时，会使用这个特性通过 `done` 通道通知 main 协程。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 使用 `jobs` 发送 3 个任务到工作协程中，然后关闭 `jobs`。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 使用前面学到的[通道同步](channel-synchronization)方法等待任务结束。
	<-done
}

// 多生产者多消费者模型
func consumer(cname string, ch chan int) {

	//可以循环 for i := range ch 来不断从 channel 接收值，直到它被关闭。

	for i := range ch {
		fmt.Println("consumer-----------", cname, ":", i)
	}
	fmt.Println("ch closed.")
}

func producer(pname string, ch chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("producer--", pname, ":", i)
		ch <- i
	}
}

func MultiProducerConsumer() {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	ch := make(chan int)
	go producer("生产者1", ch)
	// go producer("生产者2", ch)
	go consumer("消费者1", ch)
	go consumer("消费者2", ch)
	time.Sleep(3 * time.Second)
	close(ch)
	time.Sleep(3 * time.Second)
}

func immediateClose() {
	ch := make(chan int, 100)
	close(ch)

	// A
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("a error: ", err)
			}
		}()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
	fmt.Println("here")
}
