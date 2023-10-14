package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system/request"
	sysResponse "MEIS-server/model/system/response"
	"MEIS-server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct {
}

// 获取验证码
func (b *BaseApi) GetCaptcha(ctx *gin.Context) {
	id, b64s, oc, err := BaseController.GetCaptcha(ctx)
	fmt.Println("是不是这样！！！", err)
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

// 发送邮箱验证码
func (b *BaseApi) SendEmail(ctx *gin.Context) {
	var email request.SysReqEmailCode

	err2 := ctx.ShouldBindJSON(&email)

	fmt.Println("emailTo.ToMail", email, err2)

	err := utils.Verify(email, utils.EmailVerify)
	if err != nil {
		global.MEIS_LOGGER.Error("邮箱格式有误：", zap.Error(err))
		response.FailWithMessage("邮箱格式有误", ctx)
		return
	}

	str, err := MailerController.SendEmail(email)
	if err != nil {
		global.MEIS_LOGGER.Error("邮箱发送失败：", zap.Error(err))
		response.FailWithMessage("邮箱发送失败", ctx)
		return
	}
	response.SuccessWithMessage(fmt.Sprintf("邮箱发送成功当前开发模式直接返回验证码：%v", str), ctx)
	return

}
