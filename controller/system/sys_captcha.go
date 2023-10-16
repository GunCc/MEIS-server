package system

import (
	"MEIS-server/global"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type RedisCaptchaStore struct {
}

var ctx = context.Background()

const CAPTCHA = "meis-captcha:"

// 实现设置captcha方法
func (r RedisCaptchaStore) Set(id string, value string) error {
	timeout := global.MEIS_CONFIG.Captcha.Timeout // 缓存超时时间

	key := CAPTCHA + "id"
	return global.MEIS_REDIS.Set(ctx, key, value, time.Duration(timeout)).Err()
}

// 获取验证码的方法
func (r RedisCaptchaStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := global.MEIS_REDIS.Get(ctx, key).Result()
	if err != nil {
		global.MEIS_LOGGER.Error("验证码获取错误！错误：", zap.Error(err))
		return ""
	}
	if clear {
		//clear为true，验证通过，删除这个验证码
		err := global.MEIS_REDIS.Del(ctx, key).Err()
		if err != nil {
			global.MEIS_LOGGER.Error("验证码获取错误！错误：", zap.Error(err))
			return ""
		}
	}
	return val
}

//实现验证captcha的方法
func (r RedisCaptchaStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	fmt.Println("key:" + id + ";value:" + v + ";answer:" + answer)
	return v == answer
}
