package initialize

import (
	"MEIS-server/controller/system"
	. "MEIS-server/model/system"
	"context"

	"gorm.io/gorm"
)

const initOrderTables = system.InitOrderExternal - 1

// 初始化所有应该有的表单
type InitTables struct {
}

func init() {
	system.RegisterInit(initOrderTables, &InitTables{})
}

func (InitTables) InitializerName() string {
	return "ensure_tables_created"
}
func (e *InitTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *InitTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *InitTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		&JwtBlacklist{},
		&SysRole{},
		&SysUser{},
		&SysResourceType{},
		&SysResource{},
		&SysMenu{},
		&SysOperationRecord{},
		&SysUserRole{},
		&SysMenuRole{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// 视图 authority_menu 会被当成表来创建，引发冲突错误（更新版本的gorm似乎不会）
		// 由于 AutoMigrate() 基本无需考虑错误，因此显式忽略
	}
	return ctx, nil
}

func (e *InitTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		&JwtBlacklist{},
		&SysRole{},
		&SysUser{},
		&SysResourceType{},
		&SysResource{},
		&SysMenu{},
		&SysOperationRecord{},
		&SysUserRole{},
		&SysMenuRole{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
