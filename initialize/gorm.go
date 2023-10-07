package initialize

import (
	"MEIS-server/global"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		global.MEIS_LOGGER.Error("表初始化失败", zap.Error(err))
		os.Exit(0)
	}
	global.MEIS_LOGGER.Info("表初始化成功")
}
