package oa

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type AttendanceRouter struct {
}

func (p *AttendanceRouter) InitAttendanceRouter(Router *gin.RouterGroup) {
	attendanceRouter := Router.Group("oa").Group("attendance").Use(middleware.OperationRecord())
	attendanceRouterWithoutRecord := Router.Group("oa").Group("attendance")
	attendanceApi := api.ApiGroupApp.OAApi.AttendanceApi
	{
		attendanceRouter.POST("/remove", attendanceApi.RemoveAttendance) // 删除员工
		attendanceRouter.POST("/update", attendanceApi.UpdateAttendance) // 修改员工信息
		attendanceRouter.POST("/create", attendanceApi.CreateAttendance) // 修改员工信息
	}
	{
		attendanceRouterWithoutRecord.POST("/upload", attendanceApi.Upload)             // 上传信息
		attendanceRouterWithoutRecord.POST("/getList", attendanceApi.GetAttendanceList) // 员工列表
		attendanceRouterWithoutRecord.POST("/getInfo", attendanceApi.GetAttendanceInfo) // 获取员工信息
	}
}
