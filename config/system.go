package config

type System struct {
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
}
