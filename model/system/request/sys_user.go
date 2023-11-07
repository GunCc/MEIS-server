package request

type Register struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
	Enable   int    `json:"enable"`
	Code     string `json:"code"`
}

type Login struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}
