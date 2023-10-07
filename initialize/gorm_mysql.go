package initialize

import (
	"MEIS-server/global"
	"MEIS-server/initialize/internal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func GormMysql() *gorm.DB {
	ms := global.MEIS_CONFIG.Mysql
	if ms.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       ms.Dsn(),
		DefaultStringSize:         191,   // string 类型默认的长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.GormConfig()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(ms.MaxIdleConns)
		sqlDB.SetMaxOpenConns(ms.MaxOpenConns)
		return db
	}
}
