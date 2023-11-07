package system

import (
	"MEIS-server/global"

	uuid "github.com/satori/go.uuid"
)

// 用户
type SysUser struct {
	global.MEIS_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username string    `json:"username" gorm:"index;comment:用户名"`
	Avatar   string    `json:"avatar" gorm:"comment:用户头像;default:https://profile-avatar.csdnimg.cn/83f8a9da97aa4bdf85cd407e62eda7fc_weixin_44694641.jpg!1"`
	NickName string    `json:"nickname" gorm:"index;comment:昵称"`
	Password string    `json:"password" gorm:"password;comment:密码"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Enable   int       `json:"enable" gorm:"comment:是否被冻结"`
	Role     SysRole   `json:"role" gorm:"not null;comment:角色id;foreignKey:ID"`
}
