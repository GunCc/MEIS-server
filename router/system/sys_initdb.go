package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type SysInitDBRouter struct {
}

// 初始化数据库
func (s *SysInitDBRouter) InitDBRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := api.ApiGroupApp.SystemApi.SysApiInitDB
	{
		initRouter.POST("initdb", dbApi.InitDB)     // 创建数据库
		initRouter.POST("initcheck", dbApi.CheckDB) // 检索数据库
	}
}
