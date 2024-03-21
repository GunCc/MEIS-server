package oa

import (
	"MEIS-server/global"
	"time"
)

// 员工
type OATask struct {
	global.MEIS_MODEL
	TaskName    int       `json:"task_name" gorm:"comment:任务名称"`
	TaskDesc    int       `json:"task_desc" gorm:"comment:任务说明"`
	EndTime     time.Time `json:"end_time" gorm:"comment:预计完成时间"`
	PersonnelID uint      `json:"personnel_id" gorm:"任务关联的用户"`
	ProjectID   uint      `json:"project_id" gorm:"任务关联的项目"`
}

func (OATask) TableName() string {
	return "oa_task"
}
