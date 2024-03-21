package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
}

func (p *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("oa").Group("project").Use(middleware.OperationRecord())
	projectRouterWithoutRecord := Router.Group("oa").Group("project")
	projectApi := api.ApiGroupApp.OAApi.ProjectApi
	{
		projectRouter.POST("/remove", projectApi.RemoveProject) // 删除员工
		projectRouter.POST("/update", projectApi.UpdateProject) // 修改员工信息
		projectRouter.POST("/create", projectApi.CreateProject) // 修改员工信息
	}
	{
		projectRouterWithoutRecord.POST("/getList", projectApi.GetProjectList) // 员工列表
		projectRouterWithoutRecord.POST("/getInfo", projectApi.GetProjectInfo) // 获取员工信息
	}
}
