package main

import (
	"context"
	"fmt"
	"net"
	"time"
	"wwqdrh/handbook/sdk/middleware/load_balance"
	"wwqdrh/handbook/sdk/middleware/proxy"
	"wwqdrh/handbook/sdk/middleware/public"
	"wwqdrh/handbook/sdk/middleware/tcp_middleware"
	"wwqdrh/handbook/sdk/middleware/tcp_proxy"
)

var (
	addr = ":2002"
)

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, src net.Conn) {
	src.Write([]byte("tcpHandler"))
}

func main() {
	//基于 thrift 代理测试
	rb := load_balance.LoadBanlanceFactory(load_balance.LbWeightRoundRobin)
	rb.Add("127.0.0.1:6001", "40")

	//构建路由及设置中间件
	counter, _ := public.NewFlowCountService("local_app", time.Second)
	router := tcp_middleware.NewTcpSliceRouter()
	router.Group("/").Use(
		tcp_middleware.IpWhiteListMiddleWare(),
		tcp_middleware.FlowCountMiddleWare(counter))

	//构建回调handler
	routerHandler := tcp_middleware.NewTcpSliceRouterHandler(
		func(c *tcp_middleware.TcpSliceRouterContext) tcp_proxy.TCPHandler {
			return proxy.NewTcpLoadBalanceReverseProxy(c, rb)
		}, router)

	//启动服务
	tcpServ := tcp_proxy.TcpServer{Addr: addr, Handler: routerHandler}
	fmt.Println("Starting tcp_proxy at " + addr)
	tcpServ.ListenAndServe()
}
