package oa

import (
	"MEIS-server/global"
)

// 员工
type OATask struct {
	global.MEIS_MODEL
	TaskName    string      `json:"task_name" gorm:"comment:任务名称"`
	TaskDesc    string      `json:"task_desc" gorm:"comment:任务说明"`
	EndTime     string      `json:"end_time" gorm:"comment:预计完成时间"`
	Status      int         `json:"status" gorm:"comment:是否完成"`
	PersonnelID uint        `json:"personnel_id" gorm:"任务关联的用户"`
	OAPersonnel OAPersonnel `json:"personnel" gorm:"foreignKey:ID;comment:员工"`
	ProjectID   uint        `json:"project_id" gorm:"任务关联的项目"`
	OAProject   OAProject   `json:"project" gorm:"foreignKey:ID;comment:项目"`
}

func (OATask) TableName() string {
	return "oa_task"
}
