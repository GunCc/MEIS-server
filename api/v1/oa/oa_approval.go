package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	oaModel "MEIS-server/model/oa"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApprovalApi struct {
}

// 获取审批列表
func (u *ApprovalApi) GetApprovalList(ctx *gin.Context) {
	var info commenReq.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取审批列表参数错误", zap.Error(err))
		response.FailWithMessage("获取审批列表参数错误", ctx)
		return
	}

	list, total, err := ApprovalController.GetApprovalList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取审批列表参数错误", zap.Error(err))
		response.FailWithMessage("获取审批列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 修改某个审批
func (u *ApprovalApi) UpdateApproval(ctx *gin.Context) {
	var info oaModel.OAApproval

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取审批参数错误", zap.Error(err))
		response.FailWithMessage("获取审批参数错误", ctx)
		return
	}

	err = ApprovalController.UpdateApproval(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改审批错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)
}
