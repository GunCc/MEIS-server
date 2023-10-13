package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.SystemApi.BaseApi
	userApi := api.ApiGroupApp.SystemApi.UserApi
	{
		baseRouter.POST("/captcha", baseApi.GetCaptcha)
		baseRouter.POST("/sendEmail", baseApi.SendEmail)
		baseRouter.POST("/register", userApi.Register)
		baseRouter.POST("/login", userApi.Login)
	}

}
