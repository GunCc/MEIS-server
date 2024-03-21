package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
)

type TrainController struct {
}

// 获取员工列表
func (u *TrainController) GetTrainList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OATrain{})
	var trainList []oa.OATrain
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&trainList).Error
	return trainList, total, err
}

// 删除某个员工
func (i *TrainController) RemoveTrain(info oa.OATrain) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(oa.OATrain{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 添加某个员工
func (i *TrainController) CreateTrain(info oa.OATrain) (err error) {
	var trainFormDb oa.OATrain

	return global.MEIS_DB.Where("id = ?", info.ID).First(&trainFormDb).Create(info).Error
}

// 修改某个员工
func (i *TrainController) UpdateTrain(info oa.OATrain) (err error) {
	var trainFormDb oa.OATrain

	return global.MEIS_DB.Where("id = ?", info.ID).First(&trainFormDb).Updates(info).Error

}

// 获取员工信息
func (u *TrainController) GetTrainInfo(id int) (train oa.OATrain, err error) {
	err = global.MEIS_DB.First(&train, "id = ?", id).Error
	if err != nil {
		return train, err
	}
	return train, err
}
