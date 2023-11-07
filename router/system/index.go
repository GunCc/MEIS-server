package system

type SystemRouterGroup struct {
	BaseRouter
	UserRouter
	SysResourceRouter
	RoleRouter
}
