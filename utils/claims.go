package utils

import (
	"MEIS-server/global"
	"MEIS-server/model/system/request"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	// 从请求头中获取 x-token
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.MEIS_LOGGER.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// 获取uuid
func GetUserUuid(c *gin.Context) uuid.UUID {
	if value, exists := c.Get("clamis"); !exists {
		if cc, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cc.UUID
		}
	} else {
		waitUse := value.(*request.CustomClaims)
		return waitUse.UUID
	}
}
