package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysOperationRecordApi struct {
}

// 获取角色列表
func (u *SysOperationRecordApi) GetOperationRecordList(ctx *gin.Context) {
	var info request.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取日志列表参数错误", zap.Error(err))
		response.FailWithMessage("获取日志列表参数错误", ctx)
		return
	}

	list, total, err := SysOperationRecordController.GetSysOperationRecordInfoList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取日志列表参数错误", zap.Error(err))
		response.FailWithMessage("获取日志列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 删除
func (u *SysOperationRecordApi) RemoveOperationRecord(ctx *gin.Context) {
	var info system.SysOperationRecord
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("角色删除错误", zap.Error(err))
		response.FailWithMessage("角色删除错误", ctx)
		return
	}

	err = SysOperationRecordController.DeleteSysOperationRecord(info)
	if err != nil {
		global.MEIS_LOGGER.Error("角色删除错误", zap.Error(err))
		response.FailWithMessage("角色删除错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 批量删除
func (u *SysOperationRecordApi) RemoveOperationRecordByIds(ctx *gin.Context) {
	var IDS request.IdsReq

	err := ctx.ShouldBindJSON(&IDS)
	if err != nil {
		global.MEIS_LOGGER.Error("角色删除错误", zap.Error(err))
		response.FailWithMessage("角色删除错误", ctx)
		return
	}

	err = SysOperationRecordController.DeleteSysOperationRecordByIds(IDS)
	if err != nil {
		global.MEIS_LOGGER.Error("角色删除错误", zap.Error(err))
		response.FailWithMessage("角色删除错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}
