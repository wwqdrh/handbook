package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":1210"
)

////////////////////
// 路由函数
////////////////////

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye ,this is httpServer"))
}

// *handlers* 是 `net/http` 服务器里面的一个基本概念。
// handler 对象实现了 `http.Handler` 接口。
// 编写 handler 的常见方法是，在具有适当签名的函数上使用 `http.HandlerFunc` 适配器。
func hello(w http.ResponseWriter, req *http.Request) {

	// handler 函数有两个参数，`http.ResponseWriter` 和 `http.Request`。
	// response writer 被用于写入 HTTP 响应数据，这里我们简单的返回 "hello\n"。
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// 这个 handler 稍微复杂一点，
	// 我们需要读取的 HTTP 请求 header 中的所有内容，并将他们输出至 response body。
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

////////////////////
// 启动服务
////////////////////

func main() {
	Server()
	// Server2()
}

func Server() {
	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/bye", sayBye)
	// 创建服务器
	server := &http.Server{
		Addr:         Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	// 监听端口并提供服务
	log.Println("Starting httpserver at " + Addr)
	log.Fatal(server.ListenAndServe())
}

func Server2() {

	// 使用 `http.HandleFunc` 函数，可以方便的将我们的 handler 注册到服务器路由。
	// 它是 `net/http` 包中的默认路由，接受一个函数作为参数。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最后，我们调用 `ListenAndServe` 并带上端口和 handler。
	// nil 表示使用我们刚刚设置的默认路由器。
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", addr),
		Handler: http,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("api exit...")
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 在此阻塞
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("server shutdown error")
	}
}

////////////////////
// 文件相关
////////////////////
func upload(w http.ResponseWriter, r *http.Request) {
	file, head, err := r.FormFile("my_file")
	if err != nil {
		fmt.Sprintln(err)
		fmt.Fprintln(w, err)

		return
	}

	localFileDir := "/tmp/upload/"
	err = os.MkdirAll(localFileDir, 0777)
	if err != nil {
		fmt.Sprintln(err)
		fmt.Fprintln(w, err)

		return
	}

	localFilePath := localFileDir + head.Filename

	localFile, err := os.Create(localFilePath)
	if err != nil {
		fmt.Sprintln(err)
		fmt.Fprintln(w, err)

		return
	}
	defer localFile.Close()

	io.Copy(localFile, file)
	fmt.Fprintln(w, localFilePath)

}
