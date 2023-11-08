package middleware

import (
	"MEIS-server/controller"
	"MEIS-server/model/commen/response"
	"MEIS-server/utils"

	"github.com/gin-gonic/gin"
)

var JWTController = controller.ControllerGroupApp.SystemControllerGroup.JWTController

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Result(response.Unauthorized, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.Result(response.Forbidden, gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.Result(response.Forbidden, gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		// 	dr, _ := utils.ParseDuration(global.MEIS_CONFIG.JWT.ExpiresTime)
		// 	claims.ExpiresAt = time.Now().Add(dr).Unix()
		// 	newToken, _ := j.CreateTokenByOldToken(token, *claims)
		// 	newClaims, _ := j.ParseToken(newToken)
		// 	c.Header("new-token", newToken)
		// 	c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		// 	if global.MEIS_CONFIG.System.UseMultipoint {
		// 		RedisJwtToken, err := JWTController.GetJWTRedis(newClaims.Username)
		// 		if err != nil {
		// 			global.MEIS_LOGGER.Error("get redis jwt failed", zap.Error(err))
		// 		} else { // 当之前的取成功时才进行拉黑操作
		// 			// _ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
		// 		}
		// 		// 无论如何都要记录当前的活跃状态
		// 		_ = JWTController.SetJWTRedis(newToken, newClaims.Username)
		// 	}
		// }
		c.Set("claims", claims)
		c.Next()
	}
}
