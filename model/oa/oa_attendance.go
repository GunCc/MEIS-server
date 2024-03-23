package oa

import (
	"MEIS-server/global"
)

// 员工
type OAAttendance struct {
	global.MEIS_MODEL
	Work        string      `json:"work" gorm:"comment:应出勤天数"`
	Working     string      `json:"working" gorm:"comment:实际出勤"`
	IsGrand     int         `json:"is_grand" gorm:"是否发放"`
	PersonnelID uint        `json:"personnel_id" gorm:"comment:员工ID"`
	OAPersonnel OAPersonnel `json:"personnel" gorm:"foreignKey:ID;comment:员工"`
}

func (OAAttendance) TableName() string {
	return "oa_attendances"
}
