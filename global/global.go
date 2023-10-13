package global

import (
	"MEIS-server/config"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	gomail "gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var (
	MEIS_CONFIG config.Config
	MEIS_Viper  *viper.Viper
	MEIS_LOGGER *zap.Logger
	MEIS_DB     *gorm.DB
	MEIS_REDIS  *redis.Client

	MEIS_MAILER *gomail.Dialer

	// 黑名单
	BlackCache local_cache.Cache
)
