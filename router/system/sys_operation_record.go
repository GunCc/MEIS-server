package system

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct {
}

func (b *OperationRecordRouter) InitOperationRecordRouter(Router *gin.RouterGroup) {
	userRouterWithoutRecord := Router.Group("/operationRecord")
	operationRecordRouter := Router.Group("/operationRecord").Use(middleware.OperationRecord())
	operationRecordApi := api.ApiGroupApp.SystemApi.SysOperationRecordApi
	{
		userRouterWithoutRecord.POST("/getList", operationRecordApi.GetOperationRecordList)
	}
	{
		operationRecordRouter.POST("/remove", operationRecordApi.RemoveOperationRecord)
		operationRecordRouter.POST("/removeByids", operationRecordApi.RemoveOperationRecordByIds)
	}
}
