package strverify

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"strverify/utils/datastruct"
)

// 项目初始化
var (
	stringCache *datastruct.TrieTree
	cwd         string
)

// ServerStart 服务启动
func ServerStart(appCtx context.Context, https bool) {
	appInit()
	srv := &http.Server{
		Handler: appRouter(),
		Addr:    ":8080",
	}

	go func() {
		// 服务连接
		fmt.Println("https://localhost:8080")
		if https {
			if err := srv.ListenAndServeTLS(path.Join(cwd, "cert/server-cert.pem"), path.Join(cwd, "cert/server-key.pem")); err != nil {
				log.Printf("shutdown... %s", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil {
				log.Printf("shutdown... %s", err)
			}
		}

	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
		log.Println("Shutdown Server ...")
	case <-appCtx.Done():
		log.Println("context Shutdown Server ...")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// appInit 应用初始化 各种插件的加载 变量的申明等
func appInit() {
	cwd, _ = os.Getwd()
	stringCache = datastruct.NewTrieTree("")
}

func appRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.POST("/stringverify", StringVerify)

	return r
}

// StringVerify
// @Description 发送字符串 校验是否已经出现过
// @Tags
// @Accept multipart/form-data
// @Produce json
// @Param arrs formData []string true "需要验证的字符串数组"
// @Success 200 {object} stringVerifyResponse "ok"
// @Faliure 400 {object}
// @Router /stringverify [post]
func StringVerify(c *gin.Context) {
	var request stringVerifyRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}

	res := make([]bool, len(request.Arrs))
	for i, word := range request.Arrs {
		if stringCache.Search(word) {
			res[i] = true
		} else {
			stringCache.Insert(word)
			res[i] = false
		}
	}
	c.JSON(http.StatusOK, &stringVerifyResponse{
		Code: 2000,
		Data: res,
	})
}

type stringVerifyRequest struct {
	Arrs []string `form:"arrs" binding:"required"`
}

type stringVerifyResponse struct {
	Code int    `json:"code"`
	Data []bool `json:"data"`
}
