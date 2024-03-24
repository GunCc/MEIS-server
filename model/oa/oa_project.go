package oa

import (
	"MEIS-server/global"
)

// 项目
type OAProject struct {
	global.MEIS_MODEL
	ProjectName string `json:"project_name" gorm:"comment:项目名"`
	ProjectDesc string `json:"project_desc" gorm:"comment:项目描述"`
	EndTime     string `json:"end_time" gorm:"comment:结束时间"`
}

func (OAProject) TableName() string {
	return "oa_projects"
}
