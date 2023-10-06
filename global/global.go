package global

import (
	"MEIS-server/config"

	"github.com/spf13/viper"
)

var (
	MEIS_CONFIG config.Config
	MEIS_Viper  *viper.Viper
)
