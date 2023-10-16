package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"
	"context"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserController struct {
}

// 注册
func (u *UserController) Register(register request.Register) (err error) {

	if !errors.Is(global.MEIS_DB.Where("nick_name = ?", register.NickName).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("昵称重复")
	}

	var sys_role system.SysRole
	if errors.Is(global.MEIS_DB.Where("id = ?", register.RoleId).First(&sys_role).Error, gorm.ErrRecordNotFound); err != nil {
		return errors.New("角色不存在")
	}

	if register.Email != "" && !errors.Is(global.MEIS_DB.Where("email = ?", register.Email).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已经被注册")
	}

	s, err := global.MEIS_REDIS.Get(context.Background(), register.Email).Result()
	if err != nil {
		return err
	}
	fmt.Println("s", s)
	fmt.Println(" register.Code", utils.BcryptHash(register.Code))
	if !utils.BcryptCheck(s, register.Code) {
		return errors.New("验证码不匹配")
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
	fmt.Println("角色信息", sys_user)

	err = global.MEIS_DB.Create(&sys_user).Error

	return err
}

func (u *UserController) Login(login request.Login) (innerUser *system.SysUser, err error) {
	var sys_user system.SysUser

	if errors.Is(global.MEIS_DB.Where("email = ? or nick_name = ?", login.Account, login.Account).First(&sys_user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("查无此人")
	}

	fmt.Println("login.Password", login.Password)
	fmt.Println("sys_user.Password", sys_user.Password)
	if b := utils.BcryptCheck(sys_user.Password, login.Password); !b {
		return nil, errors.New("密码错误")
	}

	return &sys_user, err
}
