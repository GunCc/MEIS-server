package request

type Register struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
	Enable   int    `json:"enable"`
	Code     string `json:"code"`
	RoleIds  []uint `json:"role_ids" `
}

type Login struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captcha_id"`
}

// 用户和角色修改对象
type SetUserRoles struct {
	ID      uint
	RoleIds []uint `json:"role_ids"` // 角色ID
}

// 判断是否后台注册
func (r *Register) GetIsAdmin() bool {
	return len(r.RoleIds) != 0
}
