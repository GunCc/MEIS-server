package config

type Server struct {
	Addr int `yaml:"addr" mapstructure:"addr" json:"addr"`
}
