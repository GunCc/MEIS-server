package oa

import (
	"MEIS-server/global"
)

// 员工
type OASalary struct {
	global.MEIS_MODEL
	BaseSalary     int         `json:"base_salary" gorm:"comment:基本薪资"`
	WageSalary     int         `json:"wage_salary" gorm:"comment:绩效薪资"`
	PayslipSend    bool        `json:"payslip_send" gorm:"comment:工资条是否已发放"`
	SocialSecurity int         `json:"social_security" gorm:"comment:社保"`
	Personnel      OAPersonnel `json:"personnel" gorm:"薪资关联的用户"`
}

func (OASalary) TableName() string {
	return "oa_salary"
}
