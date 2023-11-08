package initialize

import (
	"MEIS-server/global"
	"MEIS-server/middleware"
	"MEIS-server/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.Use(middleware.Cors())

	// 格式不正确的处理中间件
	Router.Use(middleware.ErrorJsonNW())

	// 提供图片访问地址
	Router.StaticFS(global.MEIS_CONFIG.Local.StorePath, http.Dir(global.MEIS_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址

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

	PrivateGroup := Router.Group("")
	{
		SystemRouter.InitUserRouter(PrivateGroup)
		SystemRouter.InitResourceRouter(PrivateGroup)
		SystemRouter.InitRoleRouter(PrivateGroup)
		SystemRouter.InitMenuRouter(PrivateGroup)
	}

	return Router
}
