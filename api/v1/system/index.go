package system

import "MEIS-server/controller"

type SystemApi struct {
	BaseApi
	UserApi
}

var (
	BaseController = controller.ControllerGroupApp.SystemControllerGroup.UserController
	JWTController  = controller.ControllerGroupApp.SystemControllerGroup.JWTController
)
