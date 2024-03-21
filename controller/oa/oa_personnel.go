package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
	"errors"

	"gorm.io/gorm"
)

type PersonnelController struct {
}

// 获取员工列表
func (u *PersonnelController) GetPersonnelList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OAPersonnel{})
	var personnelList []oa.OAPersonnel
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&personnelList).Error
	return personnelList, total, err
}

// 删除某个员工
func (i *PersonnelController) RemovePersonnel(info oa.OAPersonnel) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(oa.OAPersonnel{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 添加某个员工
func (i *PersonnelController) CreatePersonnel(info oa.OAPersonnel) (err error) {
	if !errors.Is(global.MEIS_DB.Where("id = ? ", info.ID).First(&oa.OAPersonnel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("工号重复")
	}
	return global.MEIS_DB.Create(info).Error
}

// 修改某个员工
func (i *PersonnelController) UpdatePersonnel(info oa.OAPersonnel) (err error) {
	var personnelFormDb oa.OAPersonnel

	if !errors.Is(global.MEIS_DB.Where("id = ? and id != ?", info.ID, info.ID).First(&oa.OAPersonnel{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("工号重复")
	}

	return global.MEIS_DB.Where("id = ?", info.ID).First(&personnelFormDb).Updates(info).Error

}

// 获取员工信息
func (u *PersonnelController) GetPersonnelInfo(id int) (personnel oa.OAPersonnel, err error) {
	err = global.MEIS_DB.First(&personnel, "id = ?", id).Error
	if err != nil {
		return personnel, err
	}
	return personnel, err
}
