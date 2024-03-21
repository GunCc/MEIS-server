package initialize

import (
	"MEIS-server/global"
	"MEIS-server/model/oa"
	"MEIS-server/model/system"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	// 注册表
	err := db.AutoMigrate(
		&system.JwtBlacklist{},
		&system.SysRole{},
		&system.SysUser{},
		&system.SysResourceType{},
		&system.SysResource{},
		&system.SysMenu{},
		&system.SysOperationRecord{},
		&system.SysUserRole{},
		&system.SysMenuRole{},
		// oa表
		&oa.OAAttendance{},
		&oa.OAPersonnel{},
		&oa.OASalary{},
		&oa.OATask{},
		&oa.OATrain{},
		&oa.OAProject{},
	)
	if err != nil {
		global.MEIS_LOGGER.Error("表初始化失败", zap.Error(err))
		os.Exit(0)
	}
	global.MEIS_LOGGER.Info("表初始化成功")
}
