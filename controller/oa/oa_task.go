package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
)

type TaskController struct {
}

// 获取员工列表
func (u *TaskController) GetTaskList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OATask{})
	var taskList []oa.OATask
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&taskList).Error
	return taskList, total, err
}

// 删除某个员工
func (i *TaskController) RemoveTask(info oa.OATask) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(oa.OATask{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 添加某个员工
func (i *TaskController) CreateTask(info oa.OATask) (err error) {
	var taskFormDb oa.OATask
	return global.MEIS_DB.Where("id = ?", info.ID).First(&taskFormDb).Create(info).Error
}

// 修改某个员工
func (i *TaskController) UpdateTask(info oa.OATask) (err error) {
	var taskFormDb oa.OATask

	return global.MEIS_DB.Where("id = ?", info.ID).First(&taskFormDb).Updates(info).Error

}

// 获取员工信息
func (u *TaskController) GetTaskInfo(id int) (task oa.OATask, err error) {
	err = global.MEIS_DB.First(&task, "id = ?", id).Error
	if err != nil {
		return task, err
	}
	return task, err
}
