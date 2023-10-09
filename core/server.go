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

	// 初始化redis
	if global.MEIS_CONFIG.System.UseMultipoint || global.MEIS_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.MEIS_CONFIG.Server.Addr)

	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.MEIS_LOGGER.Info("启动成功，端口号：", zap.String("端口号：", address))
	global.MEIS_LOGGER.Error(s.ListenAndServe().Error())
}
