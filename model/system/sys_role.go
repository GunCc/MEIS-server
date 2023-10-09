package system

import "MEIS-server/global"

type SysRole struct {
	global.MEIS_MODEL
	Name string `gorm:"comment:角色名" json:"name"`
}
