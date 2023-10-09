package config

type Captcha struct {
	Open    int `yaml:"open" `
	Timeout int `yaml:"timeout" `
	Width   int `yaml:"width" `
	Height  int `yaml:"height" `
	KeyLong int `yaml:"key-long" `
}
