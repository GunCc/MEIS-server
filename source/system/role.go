package system

import (
	"MEIS-server/controller/system"
	sysModel "MEIS-server/model/system"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initRole struct {
}

// 13
const initOrderRole = initOrderCasbin + 1

// 自动运行
func init() {
	system.RegisterInit(initOrderUser, &initRole{})
}

// 自动创建表格
func (i *initRole) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysRole{})
}

// 检查表格是否已经创建
func (i *initRole) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysRole{})
}

// 获取表格名字 并且作为键值存在缓存中
func (i initRole) InitializerName() string {
	return sysModel.SysRole{}.TableName()
}

func (i *initRole) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysRole{
		{Name: "superadmin", Comment: "最高权限"},
		{Name: "test", Comment: "测试人员"},
		{Name: "customer", Comment: "游客"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", sysModel.SysRole{}.TableName())
	}

	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initRole) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("authority_id = ?", "8881").
		First(&sysModel.SysRole{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
