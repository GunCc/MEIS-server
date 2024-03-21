package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	oaModel "MEIS-server/model/oa"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AttendanceApi struct {
}

// 获取考勤列表
func (u *AttendanceApi) GetAttendanceList(ctx *gin.Context) {
	var info commenReq.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤列表参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤列表参数错误", ctx)
		return
	}

	list, total, err := AttendanceController.GetAttendanceList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤列表参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 获取考勤信息
func (u *AttendanceApi) GetAttendanceInfo(ctx *gin.Context) {
	var info commenReq.GetById
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	user, err := AttendanceController.GetAttendanceInfo(info.ID)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(user, "数据获取成功", ctx)
}

// 删除某个考勤
func (u *AttendanceApi) RemoveAttendance(ctx *gin.Context) {
	var info oaModel.OAAttendance
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	err = AttendanceController.RemoveAttendance(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 新增某个考勤
func (u *AttendanceApi) CreateAttendance(ctx *gin.Context) {
	var info oaModel.OAAttendance

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	err = AttendanceController.CreateAttendance(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改考勤错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)

}

// 修改某个考勤
func (u *AttendanceApi) UpdateAttendance(ctx *gin.Context) {
	var info oaModel.OAAttendance

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取考勤参数错误", zap.Error(err))
		response.FailWithMessage("获取考勤参数错误", ctx)
		return
	}

	err = AttendanceController.UpdateAttendance(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改考勤错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)
}
