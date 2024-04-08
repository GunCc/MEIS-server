package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type ApprovalRouter struct {
}

func (p *ApprovalRouter) InitApprovalRouter(Router *gin.RouterGroup) {
	approvalRouter := Router.Group("oa").Group("approval").Use(middleware.OperationRecord())
	approvalRouterWithoutRecord := Router.Group("oa").Group("approval")
	approvalApi := api.ApiGroupApp.OAApi.ApprovalApi
	{
		approvalRouter.POST("/update", approvalApi.UpdateApproval) // 修改审批信息
	}
	{
		approvalRouterWithoutRecord.POST("/getList", approvalApi.GetApprovalList) // 审批列表
	}
}
