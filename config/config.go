package config

type Config struct {
	Server   Server   `yaml:"server" mapstructure:"server" json:"server"`
	AutoCode AutoCode `yaml:"autocode" mapstructure:"autocode" json:"autocode"`
	Zap      Zap      `yaml:"zap" mapstructure:"zap" json:"zap"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Captcha  Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
