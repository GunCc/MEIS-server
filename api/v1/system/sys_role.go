package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system"
	"MEIS-server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysRoleApi struct {
}

func (U *SysRoleApi) CreateRole(ctx *gin.Context) {
	var role system.SysRole

	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		global.MEIS_LOGGER.Error("创建角色信息有误", zap.Error(err))
		response.Fail(ctx)
		return
	}

	err = utils.Verify(role, utils.RoleVerify)
	if err != nil {
		global.MEIS_LOGGER.Error("创建角色报错", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = RoleController.CreateRole(role)
	if err != nil {
		global.MEIS_LOGGER.Error("创建失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.MEIS_LOGGER.Info("创建成功")
	response.SuccessWithMessage("创建成功", ctx)
}

// 获取角色列表
func (u *SysRoleApi) GetRoleList(ctx *gin.Context) {
	var info request.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取角色列表参数错误", zap.Error(err))
		response.FailWithMessage("获取角色列表参数错误", ctx)
		return
	}

	list, total, err := RoleController.GetRoleList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取角色列表参数错误", zap.Error(err))
		response.FailWithMessage("获取角色列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 编辑
func (u *SysRoleApi) UpdateRole(ctx *gin.Context) {
	var info system.SysRole
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("角色编辑错误", zap.Error(err))
		response.FailWithMessage("角色编辑错误", ctx)
		return
	}

	err = RoleController.UpdateRole(info)
	if err != nil {
		global.MEIS_LOGGER.Error("角色编辑错误", zap.Error(err))
		response.FailWithMessage("角色编辑错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 删除
func (u *SysRoleApi) RemoveRole(ctx *gin.Context) {
	var info system.SysRole
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("角色删除错误", zap.Error(err))
		response.FailWithMessage("角色删除错误", ctx)
		return
	}

	err = RoleController.RemoveRole(info)
	if err != nil {
		global.MEIS_LOGGER.Error("角色删除错误", zap.Error(err))
		response.FailWithMessage("角色删除错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}
