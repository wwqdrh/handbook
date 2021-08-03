package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"archbook/restapi/models"
	"archbook/restapi/pkg/logging"
	"archbook/restapi/pkg/setting"
	"archbook/restapi/pkg/utils"
	"archbook/restapi/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	utils.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("[info] strat http server listening :%d", setting.ServerSetting.HttpPort)
	server.ListenAndServe()
}
