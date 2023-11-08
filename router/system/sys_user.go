package system

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (b *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userApi := api.ApiGroupApp.SystemApi.UserApi
	{
		userRouter.POST("/getList", userApi.GetUserList)
		userRouter.POST("/remove", userApi.RemoveUser)
		userRouter.POST("/update", userApi.UpdateUser)
		userRouter.POST("/registerAdmin", userApi.RegisterUser)
	}
}
