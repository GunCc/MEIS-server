package system

import (
	"MEIS-server/controller/system"
	. "MEIS-server/model/system"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderRoleMenu = initOrderMenu + initOrderRole

type initMenuRole struct {
}

func init() {
	system.RegisterInit(initOrderRoleMenu, &initMenuRole{})
}

// 不需要做任何事情
func (i initMenuRole) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// 每次都替代
func (i initMenuRole) TableCreated(ctx context.Context) bool {
	return false
}

func (i initMenuRole) InitializerName() string {
	return "sys_menu_roles"
}

// 插入数据
func (i initMenuRole) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	roles, ok := ctx.Value(initRole{}.InitializerName()).([]SysRole)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext, "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}

	menus, ok := ctx.Value(initMenu{}.InitializerName()).([]SysMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx

	// 超级管理员
	if err = db.Model(&roles[0]).Association("SysMenu").Replace(menus); err != nil {
		return next, err
	}

	// 测试人员
	if err = db.Model(&roles[1]).Association("SysMenu").Replace(menus); err != nil {
		return next, err
	}

	// 游客
	if err = db.Model(&roles[2]).Association("SysMenu").Replace(menus[:2]); err != nil {
		return next, err
	}
	if err = db.Model(&roles[2]).Association("SysMenu").Append(menus[4:]); err != nil {
		return next, err
	}

	return next, nil
}

func (i initMenuRole) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	role := &SysRole{}
	if ret := db.Model(role).
		Where("role_id = ?", 777).Preload("SysMenu").Find(role); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(role.SysMenu) > 0
	}
	return false
}
