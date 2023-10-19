package system

import (
	"MEIS-server/controller/system"
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/commen/response"
	systemModel "MEIS-server/model/system"
	"MEIS-server/model/system/request"
	systemRes "MEIS-server/model/system/response"
	"fmt"

	"MEIS-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type UserApi struct {
}

func (U *UserApi) Register(ctx *gin.Context) {
	var register request.Register

	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册信息有误", zap.Error(err))
		response.Fail(ctx)
		return
	}

	fmt.Println("fmt", register)
	err = utils.Verify(register, utils.RegisterVerify)
	if err != nil {
		global.MEIS_LOGGER.Error("注册校验报错", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = BaseController.Register(register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.MEIS_LOGGER.Info("注册成功")
	response.SuccessWithMessage("注册成功", ctx)
	return
}

// 登录
func (u *UserApi) Login(ctx *gin.Context) {
	var login request.Login

	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		global.MEIS_LOGGER.Error("登录参数错误", zap.Error(err))
		response.FailWithMessage("登录失败", ctx)
		return
	}

	err = utils.Verify(login, utils.LoginVerify)

	if err != nil {
		global.MEIS_LOGGER.Error("错误:", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 验证码校验
	// open := global.MEIS_CONFIG.Captcha.Open       // 是否开启防爆次数
	// timeout := global.MEIS_CONFIG.Captcha.Timeout // 缓存超时时间
	// key := ctx.ClientIP()
	// v, ok := global.BlackCache.Get(key)
	// if !ok {
	// 	global.BlackCache.Set(key, 1, time.Second*time.Duration(timeout))
	// }

	// var oc bool = open == 0 || open < BaseController.InterfaceToInt(v)

	if system.Store.Verify(login.CaptchaId, login.Captcha, true) {
		user, err := BaseController.Login(login)
		if err != nil {

			global.MEIS_LOGGER.Error("错误:", zap.Error(err))
			// 验证码次数+1
			// global.BlackCache.Increment(key, 1)
			response.FailWithMessage(err.Error(), ctx)
			return
		}

		// if user.Enalble != 1 {
		// 	global.MEIS_LOGGER.Error("登陆失败! 用户被禁止登录!")
		// 	// 验证码次数+1
		// 	// global.BlackCache.Increment(key, 1)
		// 	response.FailWithMessage("用户被禁止登录", ctx)
		// 	return
		// }
		u.TokenNext(ctx, user)
		return
	}
	// 验证码次数+1
	// global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", ctx)
}

// 下发token
func (u *UserApi) TokenNext(c *gin.Context, user *systemModel.SysUser) {
	j := &utils.JWT{
		SigningKey: []byte(global.MEIS_CONFIG.JWT.SigningKey),
	}

	claims := j.CreateClaims(request.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.MEIS_LOGGER.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 单点登录
	if !global.MEIS_CONFIG.System.UseMultipoint {
		response.SuccessWithDetailed(systemRes.UserLoginAfter{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	// 多点登录 需要用到redis
	if redisJWT, err := JWTController.GetJWTRedis(user.Username); err == redis.Nil {
		if err = JWTController.SetJWTRedis(token, user.Username); err != nil {
			global.MEIS_LOGGER.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.SuccessWithDetailed(systemRes.UserLoginAfter{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	} else if err != nil {
		global.MEIS_LOGGER.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var jwtBlacklist systemModel.JwtBlacklist

		jwtBlacklist.Jwt = redisJWT

		if err := JWTController.JoinInBlackList(jwtBlacklist); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}

		if err := JWTController.SetJWTRedis(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}

		response.SuccessWithDetailed(systemRes.UserLoginAfter{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}

}

// 获取用户列表 --- 管理员权限
func (u *UserApi) GetUserList(ctx *gin.Context) {
	var info commenReq.ListInfo

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取用户列表参数错误", zap.Error(err))
		response.FailWithMessage("获取用户列表参数错误", ctx)
		return
	}

	list, total, err := UserController.GetUserList(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取用户列表参数错误", zap.Error(err))
		response.FailWithMessage("获取用户列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}
