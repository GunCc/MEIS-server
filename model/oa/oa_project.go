package oa

import (
	"MEIS-server/global"
	"time"
)

// 项目
type OAProject struct {
	global.MEIS_MODEL
	ProjectName          int           `json:"project_name" gorm:"comment:项目名"`
	ProjectDesc          int           `json:"project_desc" gorm:"comment:项目描述"`
	EndTime              time.Time     `json:"end_time" gorm:"comment:结束时间"`
	ResponsiblePersonnel []OAPersonnel `json:"ResponsiblePersonnel" gorm:"comment:负责人"`
}

func (OAProject) TableName() string {
	return "oa_projects"
}
