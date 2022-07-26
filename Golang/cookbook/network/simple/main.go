package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(ctx *gin.Context) {
		for i := 0; i < 100; i++ {
			select {
			case closed := <-ctx.Writer.CloseNotify():
				if closed {
					fmt.Println("已经关闭")
					return
				}
			default:
				time.Sleep(1 * time.Second)
				fmt.Printf("hello %d\n", i)
			}
		}
		ctx.String(200, "ok")
	})

	srv := http.Server{
		Handler: engine,
		Addr:    ":8000",
	}

	go func() {
		srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	srv.Close()
}
