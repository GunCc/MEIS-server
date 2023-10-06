package global

import (
	"MEIS-server/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	MEIS_CONFIG config.Config
	MEIS_Viper  *viper.Viper
	MEIS_LOGGER *zap.Logger
)
