package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
	"wwqdrh/handbook/sdk/middleware/load_balance"
	"wwqdrh/handbook/sdk/middleware/proxy"
	"wwqdrh/handbook/sdk/middleware/public"
)

var addr = "127.0.0.1:2002"

// 访问限制速度
func RateLimit() {
	coreFunc := func(c *SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	sliceRouter := NewSliceRouter()
	sliceRouter.Group("/").Use(RateLimiter())
	routerHandler := NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}

func BasicRateLimit() {

	// 首先，我们将看一个基本的速率限制。
	// 假设我们想限制对收到请求的处理，我们可以通过一个渠道处理这些请求。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// `limiter` 通道每 200ms 接收一个值。
	// 这是我们任务速率限制的调度器。
	limiter := time.Tick(200 * time.Millisecond)

	// 通过在每次请求前阻塞 `limiter` 通道的一个接收，
	// 可以将频率限制为，每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
	// 我们可以通过缓冲通道来完成此任务。
	// `burstyLimiter` 通道允许最多 3 个爆发（bursts）事件。
	burstyLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发（bursts）。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200ms 我们将尝试添加一个新的值到 `burstyLimiter`中，
	// 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 现在，模拟另外 5 个传入请求。
	// 受益于 `burstyLimiter` 的爆发（bursts）能力，前 3 个请求可以快速完成。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// 熔断方案
func CircuitBreaker() {
	coreFunc := func(c *SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	public.ConfCricuitBreaker(true)
	sliceRouter := NewSliceRouter()
	sliceRouter.Group("/").Use(CircuitMW())
	routerHandler := NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}

// 访问计数
func Count() {
	coreFunc := func(c *SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	public.ConfCricuitBreaker(true)
	sliceRouter := NewSliceRouter()
	counter, _ := public.NewFlowCountService("local_app", time.Second)
	sliceRouter.Group("/").Use(FlowCountMiddleWare(counter))
	routerHandler := NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}

func RedisFlowCount() {
	coreFunc := func(c *SliceRouterContext) http.Handler {
		rs1 := "http://127.0.0.1:2003/base"
		url1, err1 := url.Parse(rs1)
		if err1 != nil {
			log.Println(err1)
		}

		rs2 := "http://127.0.0.1:2004/base"
		url2, err2 := url.Parse(rs2)
		if err2 != nil {
			log.Println(err2)
		}

		urls := []*url.URL{url1, url2}
		return proxy.NewMultipleHostsReverseProxy(c, urls)
	}
	log.Println("Starting httpserver at " + addr)

	public.ConfCricuitBreaker(true)
	sliceRouter := NewSliceRouter()
	redisCounter, _ := public.NewRedisFlowCountService("redis_app", time.Second)
	sliceRouter.Group("/").Use(RedisFlowCountMiddleWare(redisCounter))
	routerHandler := NewSliceRouterHandler(coreFunc, sliceRouter)
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}

////////////////////
// 负载均衡
////////////////////

var (
	transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, //长连接超时时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
)

func NewMultipleHostsReverseProxy(lb load_balance.LoadBalance) *httputil.ReverseProxy {
	//请求协调者
	director := func(req *http.Request) {
		nextAddr, err := lb.Get(req.RemoteAddr)
		if err != nil {
			log.Fatal("get next addr fail")
		}
		target, err := url.Parse(nextAddr)
		if err != nil {
			log.Fatal(err)
		}
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "user-agent")
		}
	}

	//更改内容
	modifyFunc := func(resp *http.Response) error {
		//请求以下命令：curl 'http://127.0.0.1:2002/error'
		if resp.StatusCode != 200 {
			//获取内容
			oldPayload, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			//追加内容
			newPayload := []byte("StatusCode error:" + string(oldPayload))
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(newPayload))
			resp.ContentLength = int64(len(newPayload))
			resp.Header.Set("Content-Length", strconv.FormatInt(int64(len(newPayload)), 10))
		}
		return nil
	}

	//错误回调 ：关闭real_server时测试，错误回调
	//范围：transport.RoundTrip发生的错误、以及ModifyResponse发生的错误
	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		//todo 如果是权重的负载则调整临时权重
		http.Error(w, "ErrorHandler error:"+err.Error(), 500)
	}

	return &httputil.ReverseProxy{Director: director, Transport: transport, ModifyResponse: modifyFunc, ErrorHandler: errFunc}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func ClientDiscovery() {
	mConf, err := load_balance.NewLoadBalanceCheckConf(
		"http://%s/base",
		map[string]string{"127.0.0.1:2003": "20", "127.0.0.1:2004": "20"})
	if err != nil {
		panic(err)
	}
	rb := load_balance.LoadBanlanceFactorWithConf(
		load_balance.LbWeightRoundRobin, mConf)
	proxy := proxy.NewLoadBalanceReverseProxy(
		&SliceRouterContext{}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

func ServerDiscovery() {
	mConf, err := load_balance.NewLoadBalanceZkConf("http://%s/base",
		"/real_server",
		[]string{"127.0.0.1:2181"},
		map[string]string{"127.0.0.1:2003": "20"})
	if err != nil {
		panic(err)
	}
	rb := load_balance.LoadBanlanceFactorWithConf(load_balance.LbWeightRoundRobin, mConf)
	proxy := proxy.NewLoadBalanceReverseProxy(&SliceRouterContext{}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
