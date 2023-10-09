package system

import "MEIS-server/controller"

type SystemApi struct {
	BaseApi
}

var (
	BaseController = controller.ControllerGroupApp.SystemControllerGroup.UserController
)
