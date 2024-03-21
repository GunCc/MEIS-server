package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	oaModel "MEIS-server/model/oa"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectApi struct {
}

// 获取项目列表
func (u *ProjectApi) GetProjectList(ctx *gin.Context) {
	var info commenReq.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目列表参数错误", zap.Error(err))
		response.FailWithMessage("获取项目列表参数错误", ctx)
		return
	}

	list, total, err := ProjectController.GetProjectList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目列表参数错误", zap.Error(err))
		response.FailWithMessage("获取项目列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 获取项目信息
func (u *ProjectApi) GetProjectInfo(ctx *gin.Context) {
	var info commenReq.GetById
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目参数错误", zap.Error(err))
		response.FailWithMessage("获取项目参数错误", ctx)
		return
	}

	user, err := ProjectController.GetProjectInfo(info.ID)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目参数错误", zap.Error(err))
		response.FailWithMessage("获取项目参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(user, "数据获取成功", ctx)
}

// 删除某个项目
func (u *ProjectApi) RemoveProject(ctx *gin.Context) {
	var info oaModel.OAProject
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目参数错误", zap.Error(err))
		response.FailWithMessage("获取项目参数错误", ctx)
		return
	}

	err = ProjectController.RemoveProject(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目参数错误", zap.Error(err))
		response.FailWithMessage("获取项目参数错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 新增某个项目
func (u *ProjectApi) CreateProject(ctx *gin.Context) {
	var info oaModel.OAProject

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目参数错误", zap.Error(err))
		response.FailWithMessage("获取项目参数错误", ctx)
		return
	}

	err = ProjectController.CreateProject(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改项目错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)

}

// 修改某个项目
func (u *ProjectApi) UpdateProject(ctx *gin.Context) {
	var info oaModel.OAProject

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取项目参数错误", zap.Error(err))
		response.FailWithMessage("获取项目参数错误", ctx)
		return
	}

	err = ProjectController.UpdateProject(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改项目错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)
}
