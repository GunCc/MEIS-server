package system

import (
	"MEIS-server/global"
)

type SysMenu struct {
	global.MEIS_MODEL
	Name      string    `json:"name" gorm:"comment:名称"`
	Path      string    `json:"path" gorm:"comment:路径"`
	Component string    `json:"component" gorm:"comment:映射组件"`
	Hidden    bool      `json:"hidden" gorm:"comment:是否隐藏;default:true"`
	ParentId  uint      `json:"p_id" gorm:"comment:父级路由id;default:0"`
	Redirect  string    `json:"redirect" gorm:"comment:重定向"`
	Children  []SysMenu `json:"children" gorm:"-"`
	Meta      `json:"meta" gorm:"embedded;comment:附加属性"`
}

func (s SysMenu) TableName() string {
	return "sys_menu"
}

type SysMenuRole struct {
	RoleId uint `json:"role_id" gorm:"comment:角色ID;column:sys_role_id"`
	MenuId uint `json:"menu_id" gorm:"comment:菜单ID;column:sys_menu_id"`
}

func (s SysMenuRole) TableName() string {
	return "sys_menu_role"
}

type Meta struct {
	Sort      string `json:"sort" gorm:"comment:排序;default:'50'"`
	Title     string `json:"title" gorm:"菜单名"`
	KeepAlive bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	Icon      string `json:"icon" gorm:"comment:菜单图标"`
}
