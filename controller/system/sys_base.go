package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
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

	if emailTo.ToMail != "" && !errors.Is(global.MEIS_DB.Where("email = ?", emailTo.ToMail).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return "发送失败", errors.New("邮箱已经被注册")
	}

	// sc, err := global.MEIS_MAILER.Dial()
	if err != nil {
		return "发送失败", err
	}

	code := u.GenRandomCode(6)
	// 生成验证码 直接顶替之前的验证码并且加密
	global.MEIS_REDIS.Set(context.Background(), emailTo.ToMail, utils.BcryptHash(code), time.Duration(global.MEIS_CONFIG.Email.ExpiresTime*1000))

	// 发送邮箱
	m := gomail.NewMessage()
	m.SetHeader("From", global.MEIS_CONFIG.Email.Account)
	m.SetHeader("To", emailTo.ToMail)
	m.SetHeader("Subject", "邮箱验证码")
	m.SetBody("text/html", fmt.Sprintf("邮箱发送成功,验证码：%v", code))
	// if err = gomail.Send(sc, m); err != nil {
	// 	return "发送失败", err
	// }
	return fmt.Sprintf("邮箱发送成功,验证码：%v", code), nil
}

const (
	words         = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6位表示字母索引
	letterIdxMask = 1<<letterIdxBits - 1 // 所有1位，与letterIdxDits一样多
	letterIdxMax  = 63 / letterIdxBits   // 适合63位的字母索引数
)

var src = rand.NewSource(time.Now().UnixNano())

// 生成随机验证码
func (u *UserController) GenRandomCode(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(words) {
			b[i] = words[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
