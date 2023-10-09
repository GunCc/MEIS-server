package controller

import "MEIS-server/controller/system"

type ControllerGroup struct {
	SystemControllerGroup system.SystemControllerGroup
}

var ControllerGroupApp = new(ControllerGroup)
