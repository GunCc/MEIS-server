package system

type SystemControllerGroup struct {
	UserController
	JWTController
	MailerController
	ResourceController
	RoleController
	MenuController
}
