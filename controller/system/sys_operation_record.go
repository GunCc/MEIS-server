package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
	"MEIS-server/model/system"
)

type SysOperationRecordController struct {
}

// 创建记录
func (or *SysOperationRecordController) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.MEIS_DB.Create(&sysOperationRecord).Error
	return err
}

// 删除记录
func (or *SysOperationRecordController) DeleteSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.MEIS_DB.Delete(&sysOperationRecord).Error
	return err
}

// 批量删除记录
func (or *SysOperationRecordController) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.MEIS_DB.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// 获取操作记录列表
func (or *SysOperationRecordController) GetSysOperationRecordInfoList(info request.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MEIS_DB.Model(&system.SysOperationRecord{})
	var sysOperationRecords []system.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.Method != "" {
	// 	db = db.Where("method = ?", info.Method)
	// }
	// if info.Path != "" {
	// 	db = db.Where("path LIKE ?", "%"+info.Path+"%")
	// }
	// if info.Status != 0 {
	// 	db = db.Where("status = ?", info.Status)
	// }
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return sysOperationRecords, total, err
}
