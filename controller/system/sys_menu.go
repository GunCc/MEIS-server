package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"errors"
	"fmt"
	"strconv"

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
	upDateMap["sort"] = menu.Sort
	if !errors.Is(global.MEIS_DB.Where("name = ? and id != ?", menu.Name, menu.ID).First(&system.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单名重复")
	}
	err = global.MEIS_DB.Model(system.SysMenu{}).Where("id = ?", menu.ID).Updates(upDateMap).Error
	return err
}

// 删除
func (u *MenuController) RemoveMenu(menu system.SysMenu) (err error) {
	err = global.MEIS_DB.Model(system.SysMenu{}).Where("id = ?", menu.ID).Delete(&menu).Error
	return err
}

// 获取菜单列表
func (u *MenuController) GetMenuList() (list interface{}, err error) {
	var menuList []system.SysMenu
	treeMap, err := u.GetAllMenuMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = u.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, err
}

// 获取整棵树
func (m *MenuController) GetAllMenuMap() (treeMap map[string][]system.SysMenu, err error) {
	var allMenu []system.SysMenu
	treeMap = make(map[string][]system.SysMenu)
	err = global.MEIS_DB.Order("sort").Find(&allMenu).Error
	for _, v := range allMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	fmt.Println("treeMap", treeMap)
	return treeMap, err
}

// 获取子节点
func (m *MenuController) getBaseChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
