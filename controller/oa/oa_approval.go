package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
)

type ApprovalController struct {
}

var NewApprovalController = new(ApprovalController)

// 获取审批列表
func (u *ApprovalController) GetApprovalList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OAApproval{})
	var approvalList []oa.OAApproval
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&approvalList).Error
	// for key, v := range approvalList {
	// 	personnel, err := NewPersonnelController.GetPersonnelInfo(int(v.PersonnelID))
	// 	if err == nil {
	// 		approvalList[key].OAPersonnel = personnel
	// 	}
	// }
	return approvalList, total, err
}

// 添加审批
func (i *ApprovalController) CreateApproval(info oa.OAApproval) (err error) {
	return global.MEIS_DB.Create(&info).Error
}

// 修改审批
func (i *ApprovalController) UpdateApproval(info oa.OAApproval) (err error) {
	return global.MEIS_DB.Where("id = ?", info.ID).Updates(&info).Error
}
