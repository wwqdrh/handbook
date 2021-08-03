package runtime

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerMux interface {
	http.Handler
}

func StartServer() {
	// 初始化global全局资源

	// 初始化gin
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
}

func StopServer() {

}
