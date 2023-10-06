package main

import (
	"MEIS-server/core"
	"MEIS-server/global"
)

func main() {

	global.MEIS_Viper = core.Viper()

	core.RunServer()
}
