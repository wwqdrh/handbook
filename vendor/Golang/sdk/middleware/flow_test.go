package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"testing"
	"time"
	"wwqdrh/handbook/sdk/middleware/load_balance"

	"github.com/afex/hystrix-go/hystrix"
	"golang.org/x/time/rate"
)

func TestRateLimiter(t *testing.T) {
	l := rate.NewLimiter(1, 5)
	log.Println(l.Limit(), l.Burst())
	for i := 0; i < 100; i++ {
		//阻塞等待直到，取到一个token
		log.Println("before Wait")
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait err:" + err.Error())
		}
		log.Println("after Wait")

		//返回需要等待多久才有新的token,这样就可以等待指定时间执行任务
		r := l.Reserve()
		log.Println("reserve Delay:", r.Delay())

		//判断当前是否可以取到token
		a := l.Allow()
		log.Println("Allow:", a)
	}
}

func TestBreaker(t *testing.T) {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8074", hystrixStreamHandler)

	hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
		Timeout:                1000, // 单次请求 超时时间
		MaxConcurrentRequests:  1,    // 最大并发量
		SleepWindow:            5000, // 熔断后多久去尝试服务是否可用
		RequestVolumeThreshold: 1,    // 验证熔断的 请求数量, 10秒内采样
		ErrorPercentThreshold:  1,    // 验证熔断的 错误百分比
	})

	for i := 0; i < 10000; i++ {
		//异步调用使用 hystrix.Go
		err := hystrix.Do("aaa", func() error {
			//test case 1 并发测试
			if i == 0 {
				return errors.New("service error")
			}
			//test case 2 超时测试
			//time.Sleep(2 * time.Second)
			log.Println("do services")
			return nil
		}, nil)
		if err != nil {
			log.Println("hystrix err:" + err.Error())
			time.Sleep(1 * time.Second)
			log.Println("sleep 1 second")
		}
	}
	time.Sleep(100 * time.Second)
}

func TestLoadBalance(t *testing.T) {
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	if err := rb.Add("http://127.0.0.1:2003/base", "10"); err != nil {
		log.Println(err)
	}
	if err := rb.Add("http://127.0.0.1:2004/base", "20"); err != nil {
		log.Println(err)
	}
	proxy := NewMultipleHostsReverseProxy(rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
