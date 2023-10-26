package system

import "MEIS-server/global"

type SysResource struct {
	global.MEIS_MODEL
	Url  string `json:"url" gorm:"comment:路径"`
	Name string `json:"name" gorm:"comment:名称"`
	Tag  string `json:"tag" gorm:"comment:标签"`
	Key  string `json:"key" gorm:"comment:编号"`
}
