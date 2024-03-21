package controller

import (
	"MEIS-server/controller/oa"
	"MEIS-server/controller/system"
)

type ControllerGroup struct {
	SystemControllerGroup system.SystemControllerGroup
	OAControllerGroup     oa.OAControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)
