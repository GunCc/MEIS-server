package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	oaModel "MEIS-server/model/oa"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SalaryApi struct {
}

// 获取薪资列表
func (u *SalaryApi) GetSalaryList(ctx *gin.Context) {
	var info commenReq.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资列表参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资列表参数错误", ctx)
		return
	}

	list, total, err := SalaryController.GetSalaryList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资列表参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// 获取薪资信息
func (u *SalaryApi) GetSalaryInfo(ctx *gin.Context) {
	var info commenReq.GetById
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}

	user, err := SalaryController.GetSalaryInfo(info.ID)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(user, "数据获取成功", ctx)
}

// 删除某个薪资
func (u *SalaryApi) RemoveSalary(ctx *gin.Context) {
	var info oaModel.OASalary
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}

	err = SalaryController.RemoveSalary(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 新增某个薪资
func (u *SalaryApi) CreateSalary(ctx *gin.Context) {
	var info oaModel.OASalary
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}

	err = SalaryController.CreateSalary(info)
	if err != nil {
		global.MEIS_LOGGER.Error("新增薪资错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("新增成功", ctx)
}

// 发放薪资某个薪资
func (u *SalaryApi) SendSalary(ctx *gin.Context) {
	var info oaModel.OASalary

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}

	msg, err := SalaryController.SendSalary(info)

	if err != nil {
		global.MEIS_LOGGER.Error("发放薪资错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	approval := oaModel.OAApproval{
		ApprovalTypeID: int(info.ID),
		ApprovalType:   oaModel.APPROVAL_SALARY,
		IsPast:         0,
	}

	// 生成审批
	ApprovalController.CreateApproval(approval)

	info.PayslipSend = true
	err = SalaryController.UpdateSalary(info)
	if err != nil {
		global.MEIS_LOGGER.Error("发放薪资错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage(msg, ctx)
}

// 修改某个薪资
func (u *SalaryApi) UpdateSalary(ctx *gin.Context) {
	var info oaModel.OASalary

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取薪资参数错误", zap.Error(err))
		response.FailWithMessage("获取薪资参数错误", ctx)
		return
	}

	err = SalaryController.UpdateSalary(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改薪资错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)
}
