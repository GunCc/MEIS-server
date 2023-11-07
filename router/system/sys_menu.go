package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (b *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	menuApi := api.ApiGroupApp.SystemApi.UserApi
	{
		menuRouter.POST("/getList", menuApi.GetUserList)
		// userRouter.POST("/remove", userApi.Login)

	}

}
