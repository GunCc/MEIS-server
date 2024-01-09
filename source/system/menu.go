package system

import (
	"MEIS-server/controller/system"
	. "MEIS-server/model/system"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenu = initOrderRole + 1

type initMenu struct{}

func init() {
	system.RegisterInit(initOrderMenu, &initMenu{})
}

func (i initMenu) InitializerName() string {
	return SysMenu{}.TableName()
}

// 创建
func (i initMenu) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&SysMenu{},
	)
}

// 查询是否创建成功
func (i initMenu) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	m := db.Migrator()
	return m.HasTable(&SysMenu{})
}

func (i initMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []SysMenu{
		{ParentId: 0, Name: "Dashboard", Path: "/dashboard", Component: "Layout", Redirect: "/dashboard/analysis", Meta: Meta{Title: "面板", Icon: "mianban"}},
		{ParentId: 1, Name: "Analysis", Path: "/dashboard/analysis", Component: "dashboard/analysis/index.vue", Meta: Meta{Title: "分析表", Icon: "mianban"}},
		{ParentId: 0, Name: "MaterialLibrary", Path: "/materialLibrary", Component: "Layout", Redirect: "/materialLibrary/manager", Meta: Meta{Title: "素材管理", Icon: "mianban", Sort: "1"}},
		{ParentId: 3, Name: "MaterialLibraryManager", Path: "/materialLibrary/manager", Component: "materialLibrary/manager/index.vue", Meta: Meta{Title: "图片管理", Icon: "mianban"}},
		{ParentId: 0, Name: "System", Path: "/system", Component: "Layout", Redirect: "/system/user", Meta: Meta{Title: "系统", Icon: "xitong"}},
		{ParentId: 5, Name: "SystemUser", Path: "/system/user", Component: "system/user/index.vue", Meta: Meta{Title: "用户管理", Icon: "mianban"}},
		{ParentId: 5, Name: "SystemRole", Path: "/system/role", Component: "system/role/index.vue", Meta: Meta{Title: "角色管理", Icon: "mianban"}},
		{ParentId: 5, Name: "SystemMenu", Path: "/system/menu", Component: "system/menu/index.vue", Meta: Meta{Title: "菜单管理", Icon: "mianban"}},
		{ParentId: 5, Name: "SystemLogger", Path: "/system/logger", Component: "system/logger/index.vue", Meta: Meta{Title: "日志管理", Icon: "mianban"}},
	}

	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, SysMenu{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i initMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "Dashboard").First(&SysMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
