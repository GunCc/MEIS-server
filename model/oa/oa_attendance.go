package oa

import (
	"MEIS-server/global"
)

// 员工
type OAAttendance struct {
	global.MEIS_MODEL
	Work        int  `json:"work" gorm:"comment:应出勤天数"`
	Working     int  `json:"working" gorm:"comment:实际出勤"`
	IsGrand     int  `gorm:"是否发放"`
	PersonnelId uint `json:"personnel_id" gorm:"comment:员工ID"`
}

func (OAAttendance) TableName() string {
	return "oa_attendances"
}
