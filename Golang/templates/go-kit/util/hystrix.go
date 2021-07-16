package util

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	ID int
	Title string
	Price int
}

func sampleGetProduct() (Product, error) {
	r := rand.Intn(10)
	if r < 6 {
		time.Sleep(time.Second * 3)
	}
	return Product{
		ID: 101,
		Title: "Golang Hystrix test",
		Price: 12,
	}, nil
}

func sampleRecProduct() (Product, error) {
	return Product{
		ID: 999,
		Title: "推荐商品",
		Price: 88,
	}, nil
}

/**
hystrix-go
熔断器：本质是隔离远程服务请求，防止级联故障

熔断器具有三种状态、关闭：默认状态。如果请求次数异常超过设定比例，则打开熔断器
打开：当熔断器打开的时候，直接执行降级方法
半开：定期的尝试发起请求来确认系统是否恢复，如果恢复了，熔断器将转为关闭状态或者保持打开
 */


func hystrixSample() {
	rand.Seed(time.Now().UnixNano())

	configA := hystrix.CommandConfig{
		Timeout: 2000,
		MaxConcurrentRequests: 5,  // 最大并发数
	}
	hystrix.ConfigureCommand("get_prod", configA)

	for {
		// Do同步执行，hystrix.Go开启协程同步执行
		err := hystrix.Do("get_prod", func() error {
			p, _ := sampleGetProduct()
			fmt.Println(p)
			return nil
		}, func(err error) error {
			//rcp, err := RecProduct()
			//fmt.Println(rcp)
			return errors.New("No Data")
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second * 1)
	}

}

func SampleAsyncHystrix() {
	rand.Seed(time.Now().UnixNano())

	configA := hystrix.CommandConfig{
		Timeout: 2000,
		MaxConcurrentRequests: 5,  // 最大并发数
		RequestVolumeThreshold: 3, // 熔断器请求阈值，默认为20，有20个请求才进行错误百分比计算
		ErrorPercentThreshold: 20, // 默认为50，开启错误熔断的百分比
	}
	hystrix.ConfigureCommand("get_prod", configA)

	resultChan := make(chan Product, 1)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go(func() {
			wg.Add(1)
			defer wg.Done()
			errs := hystrix.Go("get_prod",
				func() error {
					p, _ := sampleGetProduct()
					resultChan <- p
					return nil
				},
				func(e error) error {
					//fmt.Println(e)
					rcp, err := sampleRecProduct()
					resultChan <- rcp
					return err
				})
			select {
				case getProd := <-resultChan:
					fmt.Println(getProd)
				case err := <-errs:
					fmt.Println(err)
			}
		})()
	}
	wg.Wait()
}