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
		// "RoleId": {
		// 	NotEmpty(),
		// },
	}

	LoginVerify = Rules{
		"Account": {
			NotEmpty(),
		},
		"Password": {
			NotEmpty(),
		},
		"Captcha": {
			NotEmpty(),
		},
		"CaptchaId": {
			NotEmpty(),
		},
	}

	EmailVerify = Rules{
		"Email": {
			NotEmpty(),
		},
	}
	ResourceTypeVerify = Rules{
		"Name": {
			NotEmpty(),
		},
	}

	RoleVerify = Rules{
		"Name": {
			NotEmpty(),
		},
	}
)
