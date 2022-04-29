package channels

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	data := make(chan int)
	go producer("生产者1", data)
	// go producer("生产者2", data)
	go consumer("消费者1", data)
	go consumer("消费者2", data)

	time.Sleep(10 * time.Second)
	close(data)
	time.Sleep(10 * time.Second)
}
