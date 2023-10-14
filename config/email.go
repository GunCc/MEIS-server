package config

type Email struct {
	Account     string `yaml:"account" mapstructure:"account" json:"account"`
	AuthCode    string `yaml:"auth-code" mapstructure:"auth-code" json:"auth-code"`
	TTL         int    `yaml:"ttl" mapstructure:"ttl" json:"ttl"`
	ExpiresTime string `yaml:"expires-time" mapstructure:"expires-time" json:"expires-time"`
}
