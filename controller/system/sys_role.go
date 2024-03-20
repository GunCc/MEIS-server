package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
	"MEIS-server/model/system"
	systemReq "MEIS-server/model/system/request"
	"errors"
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

	if !errors.Is(global.MEIS_DB.Where("name = ? and id != ?", role.Name, role.RoleId).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("昵称重复")
	}
	err = global.MEIS_DB.Model(system.SysRole{}).Where("id = ?", role.RoleId).Updates(map[string]interface{}{
		"name":       role.Name,
		"updated_at": time.Now(),
		"enable":     role.Enable,
		"comment":    role.Comment,
	}).Error

	return err
}

// 删除
func (u *RoleController) RemoveRole(role system.SysRole) error {
	return global.MEIS_DB.Transaction(func(tx *gorm.DB) error {
		txErr := global.MEIS_DB.First(&system.SysUserRole{}, "sys_role_role_id = ?", role.RoleId).Error
		if !errors.Is(txErr, gorm.ErrRecordNotFound) {
			return errors.New("此角色正在被使用无法删除")
		}

		txErr = tx.Delete(&[]system.SysMenuRole{}, "sys_role_role_id = ?", role.RoleId).Error
		if txErr != nil {
			return txErr
		}

		txErr = global.MEIS_DB.Model(system.SysRole{}).Where("id = ?", role.RoleId).Delete(&role).Error
		return txErr
	})

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
	global.MEIS_DB.Preload("SysMenu").First(&role, "role_id = ?", rm.RoleId)
	var menus []system.SysMenu
	for _, v := range rm.MenuIds {
		var menu system.SysMenu
		menu.ID = v
		menus = append(menus, menu)
	}
	err = global.MEIS_DB.Model(&role).Association("SysMenu").Replace(&menus)
	return err
}
