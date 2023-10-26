package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResourceApi struct {
}

// 上传文件
func (i *ResourceApi) UploadFile(ctx *gin.Context) {
	var file system.SysResource
	// 是否保存
	noSave := ctx.DefaultQuery("noSave", "0")
	// 接收文件 参数为文件字段
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.MEIS_LOGGER.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}
	file, err = ResourceController.UploadResource(header, noSave)
	if err != nil {
		global.MEIS_LOGGER.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", ctx)
		return
	}
	response.SuccessWithDetailed(file, "上传成功", ctx)
}
