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
		salaryRouter.POST("/remove", salaryApi.RemoveSalary) // 删除员工
		salaryRouter.POST("/update", salaryApi.UpdateSalary) // 修改员工信息
		salaryRouter.POST("/create", salaryApi.CreateSalary) // 修改员工信息
	}
	{
		salaryRouterWithoutRecord.POST("/getList", salaryApi.GetSalaryList)     // 员工列表
		salaryRouterWithoutRecord.POST("/getUserInfo", salaryApi.GetSalaryInfo) // 获取员工信息
	}
}
