package system

type SystemRouterGroup struct {
	BaseRouter
	UserRouter
	SysResourceRouter
	RoleRouter
	MenuRouter
	OperationRecordRouter
	SysInitDBRouter
}
