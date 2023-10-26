package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type SysResource struct {
}

func (s *SysResource) InitResource(Router *gin.RouterGroup) {
	resourceRouter := Router.Group("resource")
	resourceApi := api.ApiGroupApp.SystemApi.ResourceApi
	{
		resourceRouter.POST("/upload", resourceApi.UploadFile)
	}
}
