package main

import (
	"MEIS-server/core"
	"MEIS-server/global"
	"MEIS-server/initialize"

	"go.uber.org/zap"
)

func main() {

	global.MEIS_Viper = core.Viper()
	global.MEIS_LOGGER = core.Zap()
	// 替换全局日志
	zap.ReplaceGlobals(global.MEIS_LOGGER)
	// 数据库链接
	global.MEIS_DB = initialize.Gorm()

	// 链接数据库

	core.RunServer()
}
