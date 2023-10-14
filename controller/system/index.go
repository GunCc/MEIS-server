package system

type SystemControllerGroup struct {
	UserController
	JWTController
	MailerController
}
