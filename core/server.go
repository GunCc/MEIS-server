package core

import (
	"MEIS-server/global"
	"MEIS-server/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.MEIS_CONFIG.Server.Addr)

	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.MEIS_LOGGER.Info("启动成功，端口号：", zap.String("端口号：", address))
	global.MEIS_LOGGER.Error(s.ListenAndServe().Error())
}
