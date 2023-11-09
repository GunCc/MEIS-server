package system

import "MEIS-server/global"

type SysRole struct {
	global.MEIS_MODEL
	Name    string    `gorm:"comment:角色名" json:"name"`
	Comment string    `gorm:"comment:备注" json:"comment"`
	Enable  int       `json:"enable" gorm:"comment:是否被冻结"`
	SysUser []SysUser `json:"-" gorm:"many2many:sys_user_role"`
}
