package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/response"
	sysResponse "MEIS-server/model/system/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct {
}

// 获取验证码
func (b *BaseApi) GetCaptcha(ctx *gin.Context) {
	id, b64s, oc, keylong, err := BaseController.GetCaptcha(ctx)
	if err != nil {
		global.MEIS_LOGGER.Error("验证码生成失败：", zap.Error(err))
		response.FailWithMessage("验证码生成失败", ctx)
		return
	}
	response.SuccessWithData(sysResponse.SysCaptcha{
		CaptchaId:     id,
		ImagePath:     b64s,
		OpenCaptcha:   oc,
		CaptchaLength: keylong,
	}, "验证码生成成功", ctx)
	return

}
