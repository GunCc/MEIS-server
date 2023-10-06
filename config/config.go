package config

type Config struct {
	Server Server `yaml:"server" mapstructure:"server" json:"server"`
}
