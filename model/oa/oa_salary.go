package oa

import (
	"MEIS-server/global"
)

// 员工
type OASalary struct {
	global.MEIS_MODEL
	BaseSalary     string      `json:"base_salary" gorm:"comment:基本薪资"`
	WageSalary     string      `json:"wage_salary" gorm:"comment:绩效薪资"`
	SocialSecurity string      `json:"social_security" gorm:"comment:社保"`
	PayslipSend    bool        `json:"payslip_send" gorm:"comment:工资条是否已发放"`
	PersonnelID    uint        `json:"personnel_id" gorm:"薪资关联的用户"`
	OAPersonnel    OAPersonnel `json:"personnel" gorm:"foreignKey:ID;comment:员工"`
	IsSend         int         `json:"is_send" gorm:"-"`
}

func (OASalary) TableName() string {
	return "oa_salary"
}
