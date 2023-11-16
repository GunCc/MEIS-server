package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
	"MEIS-server/model/system"
	systemReq "MEIS-server/model/system/request"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RoleController struct {
}

// 创建
func (u *RoleController) CreateRole(role system.SysRole) (err error) {

	if !errors.Is(global.MEIS_DB.Where("name = ?", role.Name).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("昵称重复")
	}
	err = global.MEIS_DB.Create(&role).Error

	return err
}

// 修改
func (u *RoleController) UpdateRole(role system.SysRole) (err error) {

	if !errors.Is(global.MEIS_DB.Where("name = ? and id != ?", role.Name, role.ID).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("昵称重复")
	}
	err = global.MEIS_DB.Model(system.SysRole{}).Where("id = ?", role.ID).Updates(map[string]interface{}{
		"name":       role.Name,
		"updated_at": time.Now(),
		"enable":     role.Enable,
		"comment":    role.Comment,
	}).Error

	return err
}

// 删除
func (u *RoleController) RemoveRole(role system.SysRole) (err error) {
	err = global.MEIS_DB.Model(system.SysRole{}).Where("id = ?", role.ID).Delete(&role).Error
	return err
}

// 获取用户列表
func (u *RoleController) GetRoleList(info request.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&system.SysRole{})
	var roleList []system.SysRole
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Preload("SysMenu").Offset(offset).Find(&roleList).Error
	return roleList, total, err
}

// 绑定路由和菜单
func (u *RoleController) SetRoleMenu(rm systemReq.RoleMenus) (err error) {
	var role system.SysRole

	// 找到并且打开角色菜单
	global.MEIS_DB.Preload("SysMenu").First(&role, "id = ?", rm.RoleId)
	var menus []system.SysMenu
	for _, v := range rm.MenuIds {
		var menu system.SysMenu
		menu.ID = v
		menus = append(menus, menu)
	}
	fmt.Println("menus", role)
	fmt.Println("rm", rm)
	err = global.MEIS_DB.Model(&role).Association("SysMenu").Replace(&menus)
	return err
}
