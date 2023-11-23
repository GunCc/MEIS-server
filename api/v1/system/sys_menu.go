package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system"
	"MEIS-server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysMenuApi struct {
}

func (U *SysMenuApi) CreateMenu(ctx *gin.Context) {

	var menu system.SysMenu
	err := ctx.ShouldBindJSON(&menu)

	fmt.Println("menu", menu)
	if err != nil {
		global.MEIS_LOGGER.Error("创建菜单信息有误", zap.Error(err))
		response.Fail(ctx)
		return
	}

	err = utils.Verify(menu, utils.RoleVerify)
	if err != nil {
		global.MEIS_LOGGER.Error("创建菜单报错", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = MenuController.CreateMenu(menu)
	if err != nil {
		global.MEIS_LOGGER.Error("创建失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.MEIS_LOGGER.Info("创建成功")
	response.SuccessWithMessage("创建成功", ctx)
}

// 获取菜单列表
func (u *SysMenuApi) GetMenuList(ctx *gin.Context) {

	menus, err := MenuController.GetMenuList()
	if err != nil {
		global.MEIS_LOGGER.Error("获取目录列表参数错误", zap.Error(err))
		response.FailWithMessage("获取目录列表参数错误", ctx)
		return
	}
	if menus == nil {
		menus = []system.SysMenu{}
	}
	response.SuccessWithDetailed(response.ListRes{
		List: menus,
	}, "数据获取成功", ctx)
}

// 编辑
func (u *SysMenuApi) UpdateMenu(ctx *gin.Context) {
	var info system.SysMenu
	err := ctx.ShouldBindJSON(&info)
	fmt.Println("info", info)
	if err != nil {
		global.MEIS_LOGGER.Error("菜单编辑错误", zap.Error(err))
		response.FailWithMessage("菜单编辑错误", ctx)
		return
	}

	err = MenuController.UpdateMenu(info)
	if err != nil {
		global.MEIS_LOGGER.Error("菜单编辑错误", zap.Error(err))
		response.FailWithMessage("菜单编辑错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 删除
func (u *SysMenuApi) RemoveMenu(ctx *gin.Context) {
	var info system.SysMenu
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("菜单删除错误", zap.Error(err))
		response.FailWithMessage("菜单删除错误", ctx)
		return
	}

	err = MenuController.RemoveMenu(info)
	if err != nil {
		global.MEIS_LOGGER.Error("菜单删除错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}
