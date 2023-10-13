package config

type Config struct {
	Server   Server   `yaml:"server" mapstructure:"server" json:"server"`
	AutoCode AutoCode `yaml:"autocode" mapstructure:"autocode" json:"autocode"`
	Zap      Zap      `yaml:"zap" mapstructure:"zap" json:"zap"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Captcha  Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"email"`
}
