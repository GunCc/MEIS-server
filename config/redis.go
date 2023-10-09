package config

type Redis struct {
	Db       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
