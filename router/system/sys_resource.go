package system

import (
	"MEIS-server/api"

	"github.com/gin-gonic/gin"
)

type SysResourceRouter struct {
}

func (s *SysResourceRouter) InitResourceRouter(Router *gin.RouterGroup) {
	resourceRouter := Router.Group("resource")
	resourceApi := api.ApiGroupApp.SystemApi.ResourceApi
	{
		resourceRouter.POST("/upload", resourceApi.UploadFile)
		resourceRouter.POST("/list", resourceApi.GetFileList)
		resourceRouter.POST("/remove", resourceApi.RemoveFile)
		resourceRouter.POST("/update", resourceApi.UpdateFile)
	}
}
