package system

import (
	"MEIS-server/global"
)

type SysMenu struct {
	global.MEIS_MODEL
	Name      string    `json:"name" gorm:"comment:名称"`
	Path      string    `json:"path" gorm:"comment:路径"`
	Component string    `json:"component" gorm:"comment:映射组件"`
	Sort      string    `json:"sort" gorm:"comment:排序;default:'50'"`
	Hidden    bool      `json:"hidden" gorm:"comment:是否隐藏;default:true"`
	ParentId  uint      `json:"p_id" gorm:"comment:父级路由id;default:0"`
	Children  []SysMenu `json:"children" gorm:"-"`
}
