package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"io"
	"net/url"
	"os"
	"time"
)

func serveDire() {
	// 直连方式
	tgt, _ := url.Parse("http://localhost:8080")
	client := httptransport.NewClient("GET", tgt, GetUserInfoRequest, GetUserInfoResponse)
	getUserInfo := client.Endpoint()

	ctx := context.Background()
	res, err := getUserInfo(ctx, UserRequest{Uid: 101})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := res.(UserResponse)
	fmt.Println(userinfo.Result)
}

func serveConsul() {
	// 通过consul服务中心进行获取
	config := consulapi.DefaultConfig()
	config.Address = "172.27.92.75:8500"
	apiClient, _ := consulapi.NewClient(config)
	client := consul.NewClient(apiClient)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
	}
	{
		tags := []string{"primary"}
		instancer := consul.NewInstancer(client, logger, "userservice", tags, true)
		{
			factory := func(serviceUrl string) (endpoint.Endpoint, io.Closer, error) {
				tart, _ := url.Parse("http://" + serviceUrl)
				return httptransport.NewClient("GET", tart, GetUserInfoRequest, GetUserInfoResponse).Endpoint(), nil, nil
			}
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			endpoints, _ := endpointer.Endpoints()
			fmt.Println("服务有", len(endpoints), "条")

			//mylb := lb.NewRoundRobin(endpointer) // 轮询
			mylb := lb.NewRandom(endpointer, time.Now().UnixNano())
			getUserInfo, _ := mylb.Endpoint()
			ctx := context.Background()
			res, err := getUserInfo(ctx, UserRequest{Uid: 101})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			userinfo := res.(UserResponse)
			fmt.Println(userinfo.Result)
		}
	}
}

func main() {
	serveConsul()
}
