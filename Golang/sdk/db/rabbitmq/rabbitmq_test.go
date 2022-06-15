package rabbitmq

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSimpleQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	wait := sync.WaitGroup{}
	wait.Add(2)
	go func() {
		defer wait.Done()
		// consumer
		rabbitmq := NewRabbitMQSimple("imoocSimple")
		messages := rabbitmq.ConsumeSimple(ctx)
		cnt := 0
		for mess := range messages {
			fmt.Println(string(mess))
			cnt++
			if cnt == 5 {
				cancel()
				break
			}
		}
	}()

	go func() {
		defer wait.Done()
		// producer
		rabbitmq := NewRabbitMQSimple("imoocSimple")
		rabbitmq.PublishSimple("Hello 1!")
		rabbitmq.PublishSimple("Hello 2!")
		rabbitmq.PublishSimple("Hello 3!")
		rabbitmq.PublishSimple("Hello 4!")
		rabbitmq.PublishSimple("Hello 5!")
		fmt.Println("发送成功")
	}()

	wait.Wait()
}

// TODO ERROR
func TestFanoutQueue(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	wait := sync.WaitGroup{}
	wait.Add(5)
	for i := 1; i < 5; i++ {
		go func(i int) {
			defer wait.Done()
			// consumer
			messages := NewFanoutConsumer("", ctx)
			cnt := 0
			for mess := range messages {
				fmt.Printf("%d收到消息: %s\n", i, string(mess))
				cnt++
				if cnt == 5 {
					break
				}
			}
		}(i)
	}

	go func() {
		defer wait.Done()
		// producer
		messages := make(chan string)
		NewFanoutPublish("", messages)
		messages <- "Hello 1!"
		messages <- "Hello 2!"
		messages <- "Hello 2!"
		messages <- "Hello 4!"
		messages <- "Hello 5!"
		fmt.Println("发送成功")
		close(messages)
	}()

	wait.Wait()
	time.Sleep(3 * time.Second)
	cancel()
}

func TestFanout2(t *testing.T) {
	StartConsumerMode()
}
