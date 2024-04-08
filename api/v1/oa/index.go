package oa

import "MEIS-server/controller"

type OAApi struct {
	PersonnelApi
	AttendanceApi
	ProjectApi
	SalaryApi
	TaskApi
	TrainApi
	ApprovalApi
}

var (
	PersonnelController  = controller.ControllerGroupApp.OAControllerGroup.PersonnelController
	AttendanceController = controller.ControllerGroupApp.OAControllerGroup.AttendanceController
	ProjectController    = controller.ControllerGroupApp.OAControllerGroup.ProjectController
	TaskController       = controller.ControllerGroupApp.OAControllerGroup.TaskController
	TrainController      = controller.ControllerGroupApp.OAControllerGroup.TrainController
	SalaryController     = controller.ControllerGroupApp.OAControllerGroup.SalaryController
	ApprovalController   = controller.ControllerGroupApp.OAControllerGroup.ApprovalController
)
