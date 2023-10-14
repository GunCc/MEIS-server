package config

type Email struct {
	Account     string `yaml:"account" mapstructure:"account" json:"account"`
	AuthCode    string `aml:"auth-code" mapstructure:"auth-code" json:"auth-code"`
	TTL         int    `aml:"ttl" mapstructure:"ttl" json:"ttl"`
	ExpiresTime int    `aml:" expires-time" mapstructure:" expires-time" json:" expires-time"`
}
