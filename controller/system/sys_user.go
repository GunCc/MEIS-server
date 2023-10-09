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

	if register.Email != "" && !errors.Is(global.MEIS_DB.Where("email = ?", register.Email).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已经被注册")
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

func (u *UserController) Login(login request.Login) (innerUser *system.SysUser, err error) {
	var sys_user system.SysUser

	if errors.Is(global.MEIS_DB.Where("email = ? or nickname = ?", login.Account, login.Account).First(&sys_user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("查无此人")
	}

	if b := utils.BcryptCheck(sys_user.Password, login.Password); !b {
		return nil, errors.New("密码错误")
	}

	return &sys_user, err
}
