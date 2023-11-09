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

	_, err = BaseController.Register(register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	global.MEIS_LOGGER.Info("注册成功")
	response.SuccessWithMessage("注册成功", ctx)
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
	fmt.Println("能到这里嘛", login)
	if system.Store.Verify(login.CaptchaId, login.Captcha, true) {
		fmt.Println("通过验证码校验")
		user, err := BaseController.Login(login)
		if err != nil {

			global.MEIS_LOGGER.Error("错误:", zap.Error(err))
			// 验证码次数+1
			// global.BlackCache.Increment(key, 1)
			response.FailWithMessage(err.Error(), ctx)
			return
		}

		// if user.Enable != 1 {
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

// 获取用户列表
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

// 删除某个用户
func (u *UserApi) RemoveUser(ctx *gin.Context) {
	var info systemModel.SysUser
	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取用户参数错误", zap.Error(err))
		response.FailWithMessage("获取用户参数错误", ctx)
		return
	}

	err = UserController.RemoveUser(info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取用户参数错误", zap.Error(err))
		response.FailWithMessage("获取用户参数错误", ctx)
		return
	}
	response.SuccessWithMessage("删除成功", ctx)
}

// 修改某个用户
func (u *UserApi) UpdateUser(ctx *gin.Context) {
	var info systemModel.SysUser

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.MEIS_LOGGER.Error("获取用户参数错误", zap.Error(err))
		response.FailWithMessage("获取用户参数错误", ctx)
		return
	}

	err = UserController.UpdateUser(info)
	if err != nil {
		global.MEIS_LOGGER.Error("修改用户错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = UserController.SetUserRoles(info.ID, info.RoleIds)
	if err != nil {
		global.MEIS_LOGGER.Error("修改用户绑定角色时错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("修改成功", ctx)

}

// 后台注册用户
func (u *UserApi) RegisterUser(ctx *gin.Context) {
	var register request.Register

	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册信息有误", zap.Error(err))
		response.Fail(ctx)
		return
	}

	err = utils.Verify(register, utils.RegisterVerify)
	if err != nil {
		global.MEIS_LOGGER.Error("注册校验报错", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	newUser, err := BaseController.Register(register)
	if err != nil {
		global.MEIS_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	err = UserController.SetUserRoles(newUser.ID, register.RoleIds)
	if err != nil {
		global.MEIS_LOGGER.Error("修改用户绑定角色时错误", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.SuccessWithMessage("注册成功", ctx)
}

// 重置密码
func (b *UserApi) ResetPassword(c *gin.Context) {
	var user systemModel.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = UserController.ResetPassword(user.ID)
	if err != nil {
		global.MEIS_LOGGER.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.SuccessWithMessage("重置成功", c)
}

// 绑定用户和角色的关系
func (u *UserApi) SetUserRoles(c *gin.Context) {
	var userRoles request.SetUserRoles
	err := c.ShouldBindJSON(&userRoles)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = UserController.SetUserRoles(userRoles.ID, userRoles.RoleIds)
	if err != nil {
		global.MEIS_LOGGER.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.SuccessWithMessage("修改成功", c)

}
