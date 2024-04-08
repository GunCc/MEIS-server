package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

type SalaryController struct {
}

// 获取薪资列表
func (u *SalaryController) GetSalaryList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OASalary{})
	var salaryList []oa.OASalary
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&salaryList).Error
	for key, v := range salaryList {
		personnel, err := NewPersonnelController.GetPersonnelInfo(int(v.PersonnelID))
		if err == nil {
			salaryList[key].OAPersonnel = personnel
		}

		approval, err := NewApprovalController.GetApprovalInfo(int(v.ID), oa.APPROVAL_SALARY)
		if err == nil {
			salaryList[key].IsSend = approval.IsPast
		} else {
			salaryList[key].IsSend = 2
		}
	}
	return salaryList, total, err
}

// 删除某个薪资
func (i *SalaryController) RemoveSalary(info oa.OASalary) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(oa.OASalary{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 添加某个薪资
func (i *SalaryController) CreateSalary(info oa.OASalary) (err error) {
	return global.MEIS_DB.Model(oa.OASalary{}).Create(&info).Error
}

// 修改某个薪资
func (i *SalaryController) UpdateSalary(info oa.OASalary) (err error) {
	return global.MEIS_DB.Model(oa.OASalary{}).Where("id = ?", info.ID).Updates(map[string]interface{}{
		"base_salary":     info.BaseSalary,
		"wage_salary":     info.WageSalary,
		"social_security": info.SocialSecurity,
		"payslip_send":    info.PayslipSend,
		"updated_at":      time.Now(),
		"personnel_id":    info.PersonnelID,
	}).Error

}

// 发放工资
func (i *SalaryController) SendSalary(info oa.OASalary) (msg string, err error) {
	personnel, err := NewPersonnelController.GetPersonnelInfo(int(info.PersonnelID))
	if err != nil {
		return "", errors.New("员工信息获取错误")
	}
	attendance, err := NewAttendanceController.GetAttendanceInfoByPersonnelID(int(info.PersonnelID))
	if err != nil {
		return "", errors.New("员工考勤信息获取错误")

	}
	salary, err := i.GetSalaryInfo(int(info.ID))
	if err != nil {
		return "", errors.New("员工薪资信息获取错误")
	}
	work, err := strconv.Atoi(attendance.Work)
	if err != nil {
		return "", err
	}
	working, err := strconv.Atoi(attendance.Working)
	if err != nil {
		return "", err
	}
	base, err := strconv.Atoi(salary.BaseSalary)
	if err != nil {
		return "", err
	}
	wage, err := strconv.Atoi(salary.WageSalary)
	if err != nil {
		return "", err
	}
	ss, err := strconv.Atoi(salary.SocialSecurity)
	if err != nil {
		return "", err
	}
	res_salary := (base + wage - ss) / work * working
	fmt.Println("info", personnel.Email)

	msg = fmt.Sprintf("%v你好，您应出勤天数%v，本月出勤天数%v，基本薪资%v,绩效薪资%v,社保扣除%v,最后获得薪资%v", personnel.Name, work, working, base, wage, ss, res_salary)
	// 发送邮箱
	m := gomail.NewMessage()
	m.SetHeader("From", global.MEIS_CONFIG.Email.Account)
	m.SetHeader("To", personnel.Email)
	m.SetHeader("Subject", "薪资信息")
	// m.SetBody("text/html", )
	m.SetBody("text/html", msg)
	return msg, nil
}

// 获取薪资信息
func (u *SalaryController) GetSalaryInfo(id int) (salary oa.OASalary, err error) {
	err = global.MEIS_DB.First(&salary, "id = ?", id).Error
	if err != nil {
		return salary, err
	}
	return salary, err
}
