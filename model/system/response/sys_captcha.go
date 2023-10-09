package system

type SysCaptcha struct {
	CaptchaId     string `json:"captcha_id"`
	ImagePath     string `json:"image_path"`
	CaptchaLength int    `json:"captcha_length"`
	OpenCaptcha   bool   `json:"open_captcha"`
}
