package config

type Server struct {
	Port string `yaml:"port" mapstructure:"port" json:"port"`
}
