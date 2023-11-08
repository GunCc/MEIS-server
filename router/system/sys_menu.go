package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (b *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	menuApi := api.ApiGroupApp.SystemApi.SysMenuApi
	{
		menuRouter.POST("/getList", menuApi.GetMenuList)
		menuRouter.POST("/remove", menuApi.RemoveMenu)
		menuRouter.POST("/update", menuApi.UpdateMenu)
		menuRouter.POST("/create", menuApi.CreateMenu)
	}

}
