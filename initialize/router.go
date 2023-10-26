package initialize

import (
	"MEIS-server/middleware"
	"MEIS-server/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middleware.Cors())

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
		SystemRouter.InitUserRouter(PublicGroup)
		SystemRouter.InitResourceRouter(PublicGroup)
	}

	// PrivateGroup := Router.Group("")
	{

	}

	return Router
}
