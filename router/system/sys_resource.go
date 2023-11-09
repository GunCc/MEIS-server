package system

import (
	"MEIS-server/api"
	"MEIS-server/middleware"

	"github.com/gin-gonic/gin"
)

type SysResourceRouter struct {
}

func (s *SysResourceRouter) InitResourceRouter(Router *gin.RouterGroup) {
	resourceRouter := Router.Group("resource").Use(middleware.OperationRecord())
	resourceApi := api.ApiGroupApp.SystemApi.ResourceApi
	{
		resourceRouter.POST("/upload", resourceApi.UploadFile)
		resourceRouter.POST("/list", resourceApi.GetFileList)
		resourceRouter.POST("/remove", resourceApi.RemoveFile)
		resourceRouter.POST("/update", resourceApi.UpdateFile)

		resourceRouter.POST("/fileBindType", resourceApi.FileBindType)

		resourceRouter.POST("/getFileType", resourceApi.GetFileTypeList)
		resourceRouter.POST("/addFileType", resourceApi.AddFileType)
		resourceRouter.POST("/updateFileType", resourceApi.UpdateFileType)
		resourceRouter.POST("/removeFileType", resourceApi.DeleteFileType)
	}
}
