package core

import (
	"MEIS-server/core/internal"
	"MEIS-server/global"
	"MEIS-server/utils"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if b, _ := utils.PathExists(global.MEIS_CONFIG.Zap.Director); !b {
		fmt.Println("创建文%v", global.MEIS_CONFIG.Zap.Director)
		os.Mkdir(global.MEIS_CONFIG.Zap.Director, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))
	if global.MEIS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
