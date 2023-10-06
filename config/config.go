package config

type Config struct {
	Server   Server   `yaml:"server" mapstructure:"server" json:"server"`
	AutoCode AutoCode `yaml:"autocode" mapstructure:"autocode" json:"autocode"`
}
