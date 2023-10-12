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
	id, b64s, oc, err := BaseController.GetCaptcha(ctx)
	if err != nil {
		global.MEIS_LOGGER.Error("验证码生成失败：", zap.Error(err))
		response.FailWithMessage("验证码生成失败", ctx)
		return
	}
	response.SuccessWithDetailed(sysResponse.SysCaptcha{
		CaptchaId:     id,
		ImagePath:     b64s,
		OpenCaptcha:   oc,
		CaptchaLength: global.MEIS_CONFIG.Captcha.KeyLong,
	}, "验证码生成成功", ctx)

	return

}
