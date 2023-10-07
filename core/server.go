package core

import (
	"MEIS-server/global"
	"MEIS-server/initialize"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.MEIS_CONFIG.Server.Port)
	s := initServer(address, Router)
	global.MEIS_LOGGER.Info("启动成功，端口号：", zap.String("端口号：", global.MEIS_CONFIG.Server.Port))
	global.MEIS_LOGGER.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
