package initialize

import (
	"MEIS-server/global"

	gomail "gopkg.in/gomail.v2"
)

func Mailer() *gomail.Dialer {
	conf := global.MEIS_CONFIG.Email
	dialer := gomail.NewDialer("smtp.163.com", 465, conf.Account, conf.AuthCode)

	return dialer
}
