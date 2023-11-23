package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"MEIS-server/utils"
	"errors"

	"gorm.io/gorm"
)

type MenuController struct {
}

// 创建
func (u *MenuController) CreateMenu(menu system.SysMenu) (err error) {

	if !errors.Is(global.MEIS_DB.Where("name = ?", menu.Name).First(&system.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单名重复")
	}
	err = global.MEIS_DB.Create(&menu).Error
	return err
}

// 修改
func (u *MenuController) UpdateMenu(menu system.SysMenu) (err error) {

	upDateMap := make(map[string]interface{})
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["sort"] = menu.Meta.Sort
	upDateMap["title"] = menu.Meta.Title
	upDateMap["keep_alive"] = menu.Meta.KeepAlive
	upDateMap["icon"] = menu.Meta.Icon
	if !errors.Is(global.MEIS_DB.Where("name = ? and id != ?", menu.Name, menu.ID).First(&system.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单名重复")
	}
	err = global.MEIS_DB.Model(system.SysMenu{}).Where("id = ?", menu.ID).Updates(upDateMap).Error
	return err
}

// 删除
func (u *MenuController) RemoveMenu(menu system.SysMenu) error {

	return global.MEIS_DB.Transaction(func(tx *gorm.DB) error {
		txErr := global.MEIS_DB.First(&system.SysMenuRole{}, "sys_menu_id = ?", menu.ID).Error
		if !errors.Is(txErr, gorm.ErrRecordNotFound) {
			return errors.New("此菜单正在被使用无法删除")
		}

		txErr = tx.Delete(&[]system.SysMenuRole{}, "sys_role_id = ?", menu.ID).Error
		if txErr != nil {
			return txErr
		}

		txErr = global.MEIS_DB.Model(system.SysMenu{}).Where("id = ?", menu.ID).Delete(&menu).Error
		return txErr
	})

}

// 获取菜单列表
func (u *MenuController) GetMenuList() (list interface{}, err error) {
	var menuList []system.SysMenu
	treeMap, err := u.GetAllMenuMap()
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = u.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, err
}

// 获取整棵树
func (m *MenuController) GetAllMenuMap() (treeMap map[uint][]system.SysMenu, err error) {
	var allMenu []system.SysMenu
	treeMap = make(map[uint][]system.SysMenu)
	err = global.MEIS_DB.Order("sort").Find(&allMenu).Error
	for _, v := range allMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

// 获取子节点
func (m *MenuController) getBaseChildrenList(menu *system.SysMenu, treeMap map[uint][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// 获取用户角色默认路由
func (m *MenuController) GetRoleDefaultRouter(user *system.SysUser) (treeMap map[uint][]system.SysMenu, err error) {
	var baseMenu []system.SysMenu
	var menuIds []uint
	var roleIds []uint

	for _, v := range user.Roles {
		roleIds = append(roleIds, v.ID)
	}

	err = global.MEIS_DB.Model(system.SysMenuRole{}).Where("sys_role_id in (?)", roleIds).Pluck("sys_menu_id", &menuIds).Error
	if err != nil {
		return nil, err
	}

	menuIds = utils.RemoveRep(menuIds)

	// 获取对应的菜单
	err = global.MEIS_DB.Model(system.SysMenu{}).Where("id in (?)", menuIds).Find(&baseMenu).Error

	treeMap = make(map[uint][]system.SysMenu)

	for _, v := range baseMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	menuTreeMap := make(map[uint][]system.SysMenu)
	for id, value := range treeMap {
		var pmenu system.SysMenu
		global.MEIS_DB.Where(&system.SysMenu{}).Find(&pmenu, "id = ?", id)
		pmenu.Children = value
		menuTreeMap[pmenu.ParentId] = append(menuTreeMap[pmenu.ParentId], pmenu)
	}

	return menuTreeMap, err
}

// 获取父级菜单
// func (m *MenuController) GetParentMenu(menu system.SysMenu, menus []system.SysMenu) {
// 	var pmenu system.SysMenu
// 	if menu.ParentId != 0 {
// 		// 获取父节点
// 		global.MEIS_DB.Where(&system.SysMenu{}).Find(&pmenu, "id = ?", menu.ParentId)
// 		pmenu.Children = append(pmenu.Children, menu)
// 		m.GetParentMenu(pmenu, menus)
// 	} else {
// 		menus = append(menus, menu)
// 	}
// }
