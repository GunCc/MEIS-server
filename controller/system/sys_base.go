package system

import (
	"MEIS-server/global"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var Store base64Captcha.Store = RedisCaptchaStore{}

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
	fmt.Println("keylong", height, width)

	// 字符、公式、验证码配置
	driver := base64Captcha.NewDriverDigit(height, width, keylong, 0.7, 80)

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
