package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type SalaryRouter struct {
}

func (p *SalaryRouter) InitSalaryRouter(Router *gin.RouterGroup) {
	salaryRouter := Router.Group("oa").Group("salary").Use(middleware.OperationRecord())
	salaryRouterWithoutRecord := Router.Group("oa").Group("salary")
	salaryApi := api.ApiGroupApp.OAApi.SalaryApi
	{
		salaryRouter.POST("/remove", salaryApi.RemoveSalary)   // 删除薪资
		salaryRouter.POST("/update", salaryApi.UpdateSalary)   // 修改薪资信息
		salaryRouter.POST("/create", salaryApi.CreateSalary)   // 修改薪资信息
		salaryRouter.POST("/sendSalary", salaryApi.SendSalary) // 修改薪资信息
	}
	{
		salaryRouterWithoutRecord.POST("/getList", salaryApi.GetSalaryList) // 薪资列表
		salaryRouterWithoutRecord.POST("/getInfo", salaryApi.GetSalaryInfo) // 获取薪资信息
	}
}
