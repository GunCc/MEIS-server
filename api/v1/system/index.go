package system

import "MEIS-server/controller"

type SystemApi struct {
	BaseApi
	UserApi
	ResourceApi
	SysRoleApi
	SysMenuApi
}

var (
	BaseController     = controller.ControllerGroupApp.SystemControllerGroup.UserController
	RoleController     = controller.ControllerGroupApp.SystemControllerGroup.RoleController
	MenuController     = controller.ControllerGroupApp.SystemControllerGroup.MenuController
	UserController     = controller.ControllerGroupApp.SystemControllerGroup.UserController
	MailerController   = controller.ControllerGroupApp.SystemControllerGroup.MailerController
	JWTController      = controller.ControllerGroupApp.SystemControllerGroup.JWTController
	ResourceController = controller.ControllerGroupApp.SystemControllerGroup.ResourceController
)
