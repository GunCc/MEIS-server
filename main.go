package main

import (
	"MEIS-server/core"
	"MEIS-server/global"

	"go.uber.org/zap"
)

func main() {

	global.MEIS_Viper = core.Viper()
	global.MEIS_LOGGER = core.Zap()
	// 替换全局日志
	zap.ReplaceGlobals(global.MEIS_LOGGER)
	// 数据库链接
	// global.MEIS_DB = initialize.Gorm()

	if global.MEIS_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.MEIS_DB.DB()
		defer db.Close()
	}

	core.RunServer()
}
