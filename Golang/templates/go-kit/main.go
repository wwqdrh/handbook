package main

import (
	"_examples/go-kit/Services"
	"_examples/go-kit/util"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

func main() {
	name := flag.String("name", "", "服务名称")
	port := flag.Int("p", 0, "服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名")
	}
	if *port == 0 {
		log.Fatal("请指定端口")
	}
	fmt.Println("开始运行服务，端口:", *port)
	util.SetServiceNameAndPort(*name, *port)

	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.WithPrefix(logger, "mykit", "1.0")
		logger = kitlog.With(logger, "time", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}

	user := Services.UserService{}
	limit := rate.NewLimiter(1, 5)
	endPoint := Services.RateLimit(limit)(
		Services.UserServiceMiddleware(logger)(Services.GenUserEndpoint(user)),
	)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(Services.MyErrorEncoder),
	}

	serverHandler := httptransport.NewServer(endPoint, Services.DecodeUserRequest, Services.EncodeUserResponse, options...)

	router := mux.NewRouter()
	{
		router.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
		router.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Content-type", "application/json")
			_, _ = writer.Write([]byte(`{"status": "ok"}`))
		})
	}

	errChan := make(chan error) // 构建error类型的通道
	go (func() {
		util.RegService()
		err := http.ListenAndServe(":"+strconv.Itoa(*port), router)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	})()

	go (func() {
		sigC := make(chan os.Signal)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM) // ctrl+c、正常退出
		errChan <- fmt.Errorf("%s", <-sigC)
	})()

	getErr := <-errChan // 当前两个协程出现信号传递的时候才会执行后面的内容
	util.UnRegService()
	log.Println(getErr)
}
