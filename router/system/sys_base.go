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
	{
		baseRouter.POST("/captcha", baseApi.GetCaptcha)
	}

}
