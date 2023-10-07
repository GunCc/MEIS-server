package internal

import (
	"MEIS-server/global"
	"fmt"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// 将数据库的log 转换到zap 打印出来
func (w *writer) Printf(mes string, data ...interface{}) {
	var logZap bool
	switch global.MEIS_CONFIG.System.DbType {
	case "mysql":
		logZap = global.MEIS_CONFIG.Mysql.LogZap
	default:
		logZap = global.MEIS_CONFIG.Mysql.LogZap
	}

	if logZap {
		global.MEIS_LOGGER.Info(fmt.Sprintf(mes+"\n", data...))
	} else {
		w.Writer.Printf(mes, data...)
	}
}
