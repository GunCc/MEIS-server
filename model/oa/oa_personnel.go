package oa

import (
	"MEIS-server/global"
)

// 员工
type OAPersonnel struct {
	global.MEIS_MODEL
	Name   string `json:"name" gorm:"comment:姓名"`
	Phone  string `json:"phone" gorm:"comment:手机号"`
	Email  string `json:"email" gorm:"comment:邮箱"`
	Status int    `json:"status" gorm:"comment:员工状态1入职2离职"`
}

func (OAPersonnel) TableName() string {
	return "oa_personnels"
}
