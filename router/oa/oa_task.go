package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type TaskRouter struct {
}

func (p *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("oa").Group("task").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("oa").Group("task")
	taskApi := api.ApiGroupApp.OAApi.TaskApi
	{
		taskRouter.POST("/remove", taskApi.RemoveTask) // 删除员工
		taskRouter.POST("/update", taskApi.UpdateTask) // 修改员工信息
		taskRouter.POST("/create", taskApi.CreateTask) // 修改员工信息
	}
	{
		taskRouterWithoutRecord.POST("/getList", taskApi.GetTaskList)     // 员工列表
		taskRouterWithoutRecord.POST("/getUserInfo", taskApi.GetTaskInfo) // 获取员工信息
	}
}
