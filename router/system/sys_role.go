package system

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
}

func (b *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role").Use(middleware.OperationRecord())
	roleRouterWithoutRecord := Router.Group("role")
	roleApi := api.ApiGroupApp.SystemApi.SysRoleApi
	{
		roleRouter.POST("/remove", roleApi.RemoveRole)
		roleRouter.POST("/update", roleApi.UpdateRole)
		roleRouter.POST("/create", roleApi.CreateRole)
		roleRouter.POST("/bindMenus", roleApi.SetRoleMenus)

	}
	{
		roleRouterWithoutRecord.POST("/getList", roleApi.GetRoleList)
	}
}
