package initialize

import (
	"MEIS-server/global"
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// redis的初始化
func Redis() {
	redisConfig := global.MEIS_CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.MEIS_LOGGER.Error("redis链接失败, 错误:", zap.Error(err))
	} else {
		global.MEIS_LOGGER.Info("redis链接返回:", zap.String("pong", pong))
		global.MEIS_REDIS = client
	}
}
