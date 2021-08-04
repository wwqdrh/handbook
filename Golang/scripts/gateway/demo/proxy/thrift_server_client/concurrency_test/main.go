package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/e421083458/gateway_demo/demo/proxy/thrift_server_client/gen-go/thrift_gen"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	addr := "127.0.0.1:8401"
	processTime := int64(20)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(processTime)*time.Second)

	wg := sync.WaitGroup{}
	var totalCount int64
	var successCount int64
	var failCount int64
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()

			tSocket, err := thrift.NewTSocket(addr)
			if err != nil {
				log.Fatalln("tSocket error:", err)
			}
			transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
			transport, _ := transportFactory.GetTransport(tSocket)
			protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
			client := thrift_gen.NewFormatDataClientFactory(transport, protocolFactory)
			if err := transport.Open(); err != nil {
				log.Fatalln("Error opening:", addr)
			}
			defer transport.Close()

			for {
				select {
				case <-ctx.Done():
					fmt.Println("ctx.Done")
					return
				default:
				}
				atomic.AddInt64(&totalCount, 1)
				data := thrift_gen.Data{Text: "ping"}
				_, err = client.DoFormat(context.Background(), &data)
				if err != nil {
					atomic.AddInt64(&failCount, 1)
				} else {
					atomic.AddInt64(&successCount, 1)
				}
			}
		}(ctx)
	}
	wg.Wait()
	fmt.Printf("result qps:%v succ:%v fail:%v", totalCount/processTime, successCount, failCount)
}
