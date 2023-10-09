package initialize

import (
	"MEIS-server/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")

	// 测试
	{
		PublicGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	// 无校验权限api
	SystemRouter := router.RouterGroupApp.System
	{
		SystemRouter.InitBaseRouter(PublicGroup)
	}

	// PrivateGroup := Router.Group("")
	{

	}

	return Router
}
