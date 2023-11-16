package system

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (b *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	menuApi := api.ApiGroupApp.SystemApi.SysMenuApi
	{
		menuRouter.POST("/remove", menuApi.RemoveMenu)
		menuRouter.POST("/update", menuApi.UpdateMenu)
		menuRouter.POST("/create", menuApi.CreateMenu)
	}
	{
		menuRouterWithoutRecord.POST("/getList", menuApi.GetMenuList)
	}

}
