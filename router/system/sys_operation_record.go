package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct {
}

func (b *OperationRecordRouter) InitOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("/operationRecord")
	operationRecordApi := api.ApiGroupApp.SystemApi.SysOperationRecordApi
	{
		operationRecordRouter.POST("/getList", operationRecordApi.GetOperationRecordList)
		operationRecordRouter.POST("/remove", operationRecordApi.RemoveOperationRecord)
		operationRecordRouter.POST("/removeByids", operationRecordApi.RemoveOperationRecordByIds)
	}
}
