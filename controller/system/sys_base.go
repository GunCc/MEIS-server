package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system/request"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gopkg.in/gomail.v2"
)

var Store = base64Captcha.DefaultMemStore

// 获取验证码
func (u *UserController) GetCaptcha(ctx *gin.Context) (id, b64s string, oc bool, err error) {
	// open := global.MEIS_CONFIG.Captcha.Open // 是否打开防爆
	// timeout := global.MEIS_CONFIG.Captcha.Timeout
	// // 获取ip
	// key := ctx.ClientIP()
	// v, ok := global.BlackCache.Get(key)
	// fmt.Println("v, ok", key, v, ok)
	// if !ok {
	// 	global.BlackCache.Set(key, 1, time.Second*time.Duration(timeout))
	// }

	// if open == 0 || open < u.InterfaceToInt(v) {
	// 	oc = true
	// }

	width := global.MEIS_CONFIG.Captcha.Width
	height := global.MEIS_CONFIG.Captcha.Height
	keylong := global.MEIS_CONFIG.Captcha.KeyLong

	// 字符、公式、验证码配置
	driver := base64Captcha.NewDriverDigit(width, height, keylong, 0.7, 80)

	// 验证码构造函数
	cp := base64Captcha.NewCaptcha(driver, Store)

	id, b64s, err = cp.Generate()
	// 验证码生成
	return id, b64s, oc, err
}

// 类型转换
func (u *UserController) InterfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

// 发送验证码
func (u *UserController) SendEmail(emailTo request.SysReqEmailCode) (res string, err error) {
	sc, err := global.MEIS_MAILER.Dial()
	if err != nil {
		return "发送失败", err
	}

	// 生成验证码
	global.MEIS_REDIS.Set(context.Context(), emailTo.ToMail)

	// 发送邮箱
	m := gomail.NewMessage()
	m.SetHeader("From", "MEIS---邮箱验证码")
	m.SetHeader("To", emailTo.ToMail)
	m.SetHeader("Subject", "邮箱验证码")
	m.SetBody("text/html", fmt.Sprintf("邮箱发送成功,验证码：%v", "123"))
	if err = gomail.Send(sc, m); err != nil {
		return "发送失败", err
	}
	return fmt.Sprintf("邮箱发送成功,验证码：%v", "123"), nil
}
