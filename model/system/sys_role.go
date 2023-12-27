package system

import "MEIS-server/global"

type SysRole struct {
	global.MEIS_MODEL
	Name       string    `gorm:"comment:角色名" json:"name"`
	Comment    string    `gorm:"comment:备注" json:"comment"`
	Enable     int       `json:"enable" gorm:"comment:是否被冻结"`
	SysUser    []SysUser `json:"-" gorm:"many2many:sys_user_role"`
	SysMenu    []SysMenu `json:"menus" gorm:"many2many:sys_menu_role"`
	SysMenuIds []uint    `json:"menus_ids" gorm:"-"`
}

func (SysRole) TableName() string {
	return "sys_roles"
}
