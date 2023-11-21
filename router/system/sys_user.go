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
	userRouterWithoutRecord := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApi.UserApi
	{
		userRouter.POST("/remove", userApi.RemoveUser)           // 删除用户
		userRouter.POST("/update", userApi.UpdateUser)           // 修改用户信息
		userRouter.POST("/resetPassword", userApi.ResetPassword) // 重置密码
		userRouter.POST("/setUserRoles", userApi.SetUserRoles)   // 设置用户权限组
		userRouter.POST("/registerAdmin", userApi.RegisterUser)  // 管理端注册用户
	}
	{
		userRouterWithoutRecord.POST("/getList", userApi.GetUserList)     // 用户列表
		userRouterWithoutRecord.POST("/getUserInfo", userApi.GetUserInfo) // 用户列表
	}
}
