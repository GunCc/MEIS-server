package utils

import (
	"MEIS-server/global"
	"MEIS-server/model/system/request"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")            // 过期
	TokenNotValidYet = errors.New("Token not active yet")        // 未激活
	TokenMalformed   = errors.New("That's not even a token")     // 不是token
	TokenInvalid     = errors.New("Couldn't handle this token:") // 无法处理
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.MEIS_CONFIG.JWT.SigningKey),
	}
}

// 创建一个声明
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	// 缓冲时间
	bf, _ := ParseDuration(global.MEIS_CONFIG.JWT.BufferTime)
	// 有效时间
	ep, _ := ParseDuration(global.MEIS_CONFIG.JWT.ExpiresTime)

	clamis := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), //缓冲时间一天
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,      // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),     // 过期时间 7天  配置文件
			Issuer:    global.MEIS_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}
	return clamis
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
