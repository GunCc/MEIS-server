package utils

var (
	RegisterVerify = Rules{
		"NickName": {
			NotEmpty(),
		},
		"Password": {
			NotEmpty(),
		},
		"Email": {
			NotEmpty(),
		},
		"RoleId": {
			NotEmpty(),
		},
	}
)
