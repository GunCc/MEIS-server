package system

import "time"

type SysRole struct {
	CreatedAt  time.Time  // 创建时间
	UpdatedAt  time.Time  // 更新时间
	DeletedAt  *time.Time `sql:"index"`
	RoleId     uint       `json:"role_id" gorm:"auto_increment null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	Name       string     `gorm:"comment:角色名" json:"name"`
	Comment    string     `gorm:"comment:备注" json:"comment"`
	Enable     int        `json:"enable" gorm:"comment:是否被冻结"`
	SysUser    []SysUser  `json:"-" gorm:"many2many:sys_user_role"`
	SysMenu    []SysMenu  `json:"menus" gorm:"many2many:sys_menu_role"`
	SysMenuIds []uint     `json:"menus_ids" gorm:"-"`
}

func (SysRole) TableName() string {
	return "sys_roles"
}
