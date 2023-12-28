package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/response"
	"MEIS-server/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysApiInitDB struct {
}

// 初始化数据库
func (*SysApiInitDB) InitDB(c *gin.Context) {
	if global.MEIS_DB != nil {
		response.SuccessWithMessage("已存在数据库配置", c)
		return
	}

	var dbInfo request.InitDB

	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.MEIS_LOGGER.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := InitDBController.InitDB(dbInfo); err != nil {
		global.MEIS_LOGGER.Error("自动创建数据库失败!", zap.Error(err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.SuccessWithMessage("自动创建数据库成功", c)
}

// 检索数据库是否初始化

func (i *SysApiInitDB) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.MEIS_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.MEIS_LOGGER.Info(message)
	response.SuccessWithDetailed(gin.H{"needInit": needInit}, message, c)
}
