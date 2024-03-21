package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type TrainRouter struct {
}

func (p *TrainRouter) InitTrainRouter(Router *gin.RouterGroup) {
	trainRouter := Router.Group("oa").Group("train").Use(middleware.OperationRecord())
	trainRouterWithoutRecord := Router.Group("oa").Group("train")
	trainApi := api.ApiGroupApp.OAApi.TrainApi
	{
		trainRouter.POST("/remove", trainApi.RemoveTrain) // 删除员工
		trainRouter.POST("/update", trainApi.UpdateTrain) // 修改员工信息
		trainRouter.POST("/create", trainApi.CreateTrain) // 修改员工信息
	}
	{
		trainRouterWithoutRecord.POST("/getList", trainApi.GetTrainList) // 员工列表
		trainRouterWithoutRecord.POST("/getInfo", trainApi.GetTrainInfo) // 获取员工信息
	}
}
