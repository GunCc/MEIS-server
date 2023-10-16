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

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

const (
	words         = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6位表示字母索引
	letterIdxMask = 1<<letterIdxBits - 1 // 所有1位，与letterIdxDits一样多
	letterIdxMax  = 63 / letterIdxBits   // 适合63位的字母索引数
)

type MailerController struct {
}

// 发送验证码
func (mail *MailerController) SendEmail(emailTo request.SysReqEmailCode) (res string, err error) {

	if emailTo.ToMail != "" && !errors.Is(global.MEIS_DB.Where("email = ?", emailTo.ToMail).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return "发送失败", errors.New("邮箱已经被注册")
	}

	// sc, err := global.MEIS_MAILER.Dial()
	if err != nil {
		return "发送失败", err
	}

	code := mail.GenRandomCode(6)
	dr, err := time.ParseDuration(global.MEIS_CONFIG.Email.ExpiresTime)
	if err != nil {
		return "时间格式化错误", err
	}
	// 生成验证码 直接顶替之前的验证码并且加密
	global.MEIS_REDIS.Set(context.Background(), emailTo.ToMail, utils.BcryptHash(code), dr)

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

// 生成随机验证码
func (mail *MailerController) GenRandomCode(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
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
