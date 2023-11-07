package system

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/system"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"
	"context"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserController struct {
}

// 注册
func (u *UserController) Register(register request.Register, haveCode bool) (err error) {

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

	if haveCode {
		s, err := global.MEIS_REDIS.Get(context.Background(), register.Email).Result()
		if err != nil {
			return err
		}

		if !utils.BcryptCheck(s, register.Code) {
			return errors.New("验证码不匹配")
		}

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

// 登录
func (u *UserController) Login(login request.Login) (innerUser *system.SysUser, err error) {
	var sys_user system.SysUser

	if errors.Is(global.MEIS_DB.Where("email = ? or nick_name = ?", login.Account, login.Account).First(&sys_user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("查无此人")
	}

	if b := utils.BcryptCheck(sys_user.Password, login.Password); !b {
		return nil, errors.New("密码错误")
	}

	return &sys_user, err
}

// 获取用户列表
func (u *UserController) GetUserList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

// 删除某个用户
func (i *UserController) RemoveUser(info system.SysUser) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(system.SysUser{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 修改某个用户
func (i *UserController) UpdateUser(info system.SysUser) (err error) {
	var userFormDb system.SysUser

	if !errors.Is(global.MEIS_DB.Where("nick_name = ? and id != ?", info.NickName, info.ID).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("昵称重复")
	}

	if info.Email != "" && !errors.Is(global.MEIS_DB.Where("email = ? and id != ?", info.Email, info.ID).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已经被注册")
	}
	return global.MEIS_DB.Where("id = ?", info.ID).First(&userFormDb).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"nick_name":  info.NickName,
		"avatar":     info.Avatar,
		"email":      info.Email,
		"enable":     info.Enable,
	}).Error

}
