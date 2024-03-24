package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
	"time"
)

type TaskController struct {
}

// 获取任务列表
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
	for key, v := range taskList {
		personnel, err := NewPersonnelController.GetPersonnelInfo(int(v.PersonnelID))
		if err == nil {
			taskList[key].OAPersonnel = personnel
		}

		project, err := NewProjectController.GetProjectInfo(int(v.ProjectID))
		if err == nil {
			taskList[key].OAProject = project
		}
	}
	return taskList, total, err
}

// 删除某个任务
func (i *TaskController) RemoveTask(info oa.OATask) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(oa.OATask{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 添加某个任务
func (i *TaskController) CreateTask(info oa.OATask) (err error) {
	return global.MEIS_DB.Create(&info).Error
}

// 修改某个任务
func (i *TaskController) UpdateTask(info oa.OATask) (err error) {
	return global.MEIS_DB.Model(oa.OATask{}).Where("id = ?", info.ID).Updates(map[string]interface{}{
		"task_name":   info.TaskName,
		"task_desc":   info.TaskDesc,
		"end_time":    info.EndTime,
		"status":      info.Status,
		"updated_at":  time.Now(),
		"PersonnelID": info.PersonnelID,
		"ProjectID":   info.ProjectID,
	}).Error
}

// 获取任务信息
func (u *TaskController) GetTaskInfo(id int) (task oa.OATask, err error) {
	err = global.MEIS_DB.First(&task, "id = ?", id).Error
	if err != nil {
		return task, err
	}
	return task, err
}
