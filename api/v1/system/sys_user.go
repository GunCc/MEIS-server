package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

func (U *UserApi) Register(ctx *gin.Context) {
	var register request.Register

	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册信息有误", zap.Error(err))
		response.Fail(ctx)
		return
	}

	err = utils.Verify(register, utils.RegisterVerify)
	if err != nil {
		global.MEIS_LOGGER.Error("注册校验报错", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = BaseController.Register(register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.MEIS_LOGGER.Info("注册成功")
	response.SuccessWithMessage("注册成功", ctx)
	return
}
