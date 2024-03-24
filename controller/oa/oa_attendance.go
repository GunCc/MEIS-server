package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
)

type AttendanceController struct {
}

// 获取考勤列表
func (u *AttendanceController) GetAttendanceList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OAAttendance{})
	var attendanceList []oa.OAAttendance
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&attendanceList).Error
	for key, v := range attendanceList {
		personnel, err := NewPersonnelController.GetPersonnelInfo(int(v.PersonnelID))
		if err == nil {
			attendanceList[key].OAPersonnel = personnel
		}
	}
	return attendanceList, total, err
}

// 删除某个考勤
func (i *AttendanceController) RemoveAttendance(info oa.OAAttendance) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Delete(&oa.OAAttendance{}, info.ID).Error
	return err
}

// 添加某个考勤
func (i *AttendanceController) CreateAttendance(info oa.OAAttendance) (err error) {
	return global.MEIS_DB.Create(&info).Error
}

// 修改某个考勤
func (i *AttendanceController) UpdateAttendance(info oa.OAAttendance) (err error) {
	return global.MEIS_DB.Where("id = ?", info.ID).Updates(&info).Error
}

// 获取考勤信息
func (u *AttendanceController) GetAttendanceInfo(id int) (attendance oa.OAAttendance, err error) {
	err = global.MEIS_DB.First(&attendance, "id = ?", id).Error
	if err != nil {
		return attendance, err
	}
	return attendance, err
}
