package request

type Register struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `json:"role"`
	Enalble  int    `json:"enable"`
}
