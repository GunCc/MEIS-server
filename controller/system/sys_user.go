package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserController struct {
}

// 注册
func (u *UserController) Register(register request.Register) (err error) {

	if !errors.Is(global.MEIS_DB.Where("nickname = ?", register.NickName).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("昵称重复")
	}

	var sys_role system.SysRole
	if errors.Is(global.MEIS_DB.Where("id = ?", register.RoleId).First(&sys_role).Error, gorm.ErrRecordNotFound); err != nil {
		return errors.New("角色不存在")
	}
	sys_user := system.SysUser{
		NickName: register.NickName,
		Email:    register.Email,
	}

	sys_user.Role = sys_role

	// 附加uuid
	sys_user.UUID = uuid.NewV4()

	// 加密password
	sys_user.Password = utils.BcryptHash(register.Password)

	err = global.MEIS_DB.Create(&sys_user).Error

	return err
}
