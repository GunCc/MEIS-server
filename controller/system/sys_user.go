package system

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/system"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"
	"context"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserController struct {
}

// 注册
func (u *UserController) Register(register request.Register) (user *system.SysUser, err error) {

	if !errors.Is(global.MEIS_DB.Where("nick_name = ?", register.NickName).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("昵称重复")
	}

	if register.Email != "" && !errors.Is(global.MEIS_DB.Where("email = ?", register.Email).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("邮箱已经被注册")
	}

	fmt.Println("register.GetIsAdmin()", register.GetIsAdmin())

	// 如果是后台注册不需要验证码
	if !register.GetIsAdmin() {
		s, err := global.MEIS_REDIS.Get(context.Background(), register.Email).Result()
		if err != nil {
			return nil, err
		}

		if !utils.BcryptCheck(s, register.Code) {
			return nil, errors.New("验证码不匹配")
		}

	}

	sys_user := system.SysUser{
		NickName: register.NickName,
		Email:    register.Email,
	}

	// 附加uuid
	sys_user.UUID = uuid.NewV4()

	// 加密password
	sys_user.Password = utils.BcryptHash(register.Password)

	err = global.MEIS_DB.Create(&sys_user).Error

	return &sys_user, err
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
	err = db.Limit(limit).Offset(offset).Preload("Roles").Find(&userList).Error
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

// 重置用户密码
func (u *UserController) ResetPassword(id uint) (err error) {

	return global.MEIS_DB.Model(&system.SysUser{}).Where("id = ?", id).Update("password", utils.BcryptHash("123456")).Error
}

// 修改用户和角色关系

func (u *UserController) SetUserRoles(id uint, roleIds []uint) (err error) {
	return global.MEIS_DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&[]system.SysUserRole{}, "sys_user_id = ?", id).Error
		if err != nil {
			return err
		}

		var userRoles []system.SysUserRole
		for _, v := range roleIds {
			userRoles = append(userRoles, system.SysUserRole{
				SysUserId: id,
				SysRoleId: v,
			})
		}
		err = tx.Create(userRoles).Error
		if err != nil {
			return err
		}

		return nil
	})
}
