package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
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
	var salaryFormDb oa.OASalary
	return global.MEIS_DB.Where("id = ?", info.ID).First(&salaryFormDb).Create(info).Error
}

// 修改某个薪资
func (i *SalaryController) UpdateSalary(info oa.OASalary) (err error) {
	var salaryFormDb oa.OASalary

	return global.MEIS_DB.Where("id = ?", info.ID).First(&salaryFormDb).Updates(info).Error

}

// 获取薪资信息
func (u *SalaryController) GetSalaryInfo(id int) (salary oa.OASalary, err error) {
	err = global.MEIS_DB.First(&salary, "id = ?", id).Error
	if err != nil {
		return salary, err
	}
	return salary, err
}
