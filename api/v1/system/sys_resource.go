package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
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

// 删除某个资源文件
func (i *ResourceApi) RemoveFile(ctx *gin.Context) {
	var file system.SysResource

	err := ctx.ShouldBindJSON(&file)
	if err != nil {
		global.MEIS_LOGGER.Error("获取文件资源报错：", zap.Error(err))
		response.FailWithMessage("获取文件资源报错：", ctx)
		return
	}

	if err := ResourceController.RemoveFile(file); err != nil {
		global.MEIS_LOGGER.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)

}

// 编辑某个资源文件
func (i *ResourceApi) UpdateFile(ctx *gin.Context) {
	var file system.SysResource

	err := ctx.ShouldBindJSON(&file)
	if err != nil {
		global.MEIS_LOGGER.Error("获取文件资源报错：", zap.Error(err))
		response.FailWithMessage("获取文件资源报错：", ctx)
		return
	}

	if err := ResourceController.UpdateFile(file); err != nil {
		global.MEIS_LOGGER.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", ctx)
		return
	}
	response.SuccessWithMessage("编辑成功", ctx)

}

// 获取文件资源列表
func (i *ResourceApi) GetFileList(ctx *gin.Context) {
	var info request.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取资源列表参数错误", zap.Error(err))
		response.FailWithMessage("获取资源列表参数错误", ctx)
		return
	}

	list, total, err := ResourceController.GetResourceList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取资源列表参数错误", zap.Error(err))
		response.FailWithMessage("获取资源列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}
