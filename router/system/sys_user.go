package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (b *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApi.UserApi
	{
		userRouter.POST("/getList", userApi.GetUserList)
		// userRouter.POST("/remove", userApi.Login)

	}

}
