package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"MEIS-server/utils"
	"context"
)

type JWTController struct {
}

// 从redis中获取jwt
func (j *JWTController) GetJWTRedis(username string) (redisJWT string, err error) {
	redisJWT, err = global.MEIS_REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

// 将jwt存入redis并且设置过期时间
func (j *JWTController) SetJWTRedis(token string, username string) error {
	// 处理jwt过期时间
	dr, err := utils.ParseDuration(global.MEIS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	err = global.MEIS_REDIS.Set(context.Background(), username, token, dr).Err()
	return err
}

// 拉黑jwt
func (j *JWTController) JoinInBlackList(jwtList system.JwtBlacklist) (err error) {
	err = global.MEIS_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}
