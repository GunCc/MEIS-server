package oa

import (
	"MEIS-server/global"
	commenReq "MEIS-server/model/commen/request"
	"MEIS-server/model/oa"
	"errors"

	"gorm.io/gorm"
)

type ProjectController struct {
}

var NewProjectController = new(ProjectController)

// 获取项目列表
func (u *ProjectController) GetProjectList(info commenReq.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&oa.OAProject{})
	var projectList []oa.OAProject
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projectList).Error
	return projectList, total, err
}

// 删除某个项目
func (i *ProjectController) RemoveProject(info oa.OAProject) (err error) {
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Model(oa.OAProject{}).Where("id = ?", info.ID).Delete(&info).Error
	return err
}

// 添加某个项目
func (i *ProjectController) CreateProject(info oa.OAProject) (err error) {
	return global.MEIS_DB.Create(&info).Error
}

// 修改某个项目
func (i *ProjectController) UpdateProject(info oa.OAProject) (err error) {
	var projectFormDb oa.OAProject
	if !errors.Is(global.MEIS_DB.Where(" id == ?", info.ID).First(&oa.OAProject{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("找不到改项目")
	}
	return global.MEIS_DB.Where("id = ?", info.ID).First(&projectFormDb).Updates(&info).Error

}

// 获取项目信息
func (u *ProjectController) GetProjectInfo(id int) (project oa.OAProject, err error) {
	err = global.MEIS_DB.First(&project, "id = ?", id).Error
	if err != nil {
		return project, err
	}
	return project, err
}
