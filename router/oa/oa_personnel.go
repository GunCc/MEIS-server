package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type PersonnelRouter struct {
}

func (p *PersonnelRouter) InitPersonnelRouter(Router *gin.RouterGroup) {
	personnelRouter := Router.Group("oa").Group("personnel").Use(middleware.OperationRecord())
	personnelRouterWithoutRecord := Router.Group("oa").Group("personnel")
	personnelApi := api.ApiGroupApp.OAApi.PersonnelApi
	{
		personnelRouter.POST("/remove", personnelApi.RemovePersonnel) // 删除员工
		personnelRouter.POST("/update", personnelApi.UpdatePersonnel) // 修改员工信息
		personnelRouter.POST("/create", personnelApi.CreatePersonnel) // 修改员工信息
	}
	{
		personnelRouterWithoutRecord.POST("/getList", personnelApi.GetPersonnelList)     // 员工列表
		personnelRouterWithoutRecord.POST("/getUserInfo", personnelApi.GetPersonnelInfo) // 获取员工信息
	}
}
