package config

type Captcha struct {
	Open    int `yaml:"open" mapstructure:"open" json:"open"`
	Timeout int `yaml:"timeout" mapstructure:"timeout" json:"timeout"`
	Width   int `yaml:"width" mapstructure:"width" json:"width"`
	Height  int `yaml:"height" mapstructure:"height" json:"height"`
	KeyLong int `yaml:"key-long" mapstructure:"key-long" json:"key-long"`
}
