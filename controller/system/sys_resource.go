package system

import (
	"MEIS-server/global"
	"MEIS-server/model/commen/request"
	"MEIS-server/model/system"
	"MEIS-server/utils/upload"
	"errors"
	"mime/multipart"
	"strings"
)

type ResourceController struct {
}

// 上传资源
func (i *ResourceController) UploadResource(header *multipart.FileHeader, noSave string) (file system.SysResource, err error) {
	oss := upload.NewOOS()
	filePath, key, uploadErr := oss.UploadFile(header)
	// 如果上传失败
	if uploadErr != nil {
		panic(err)
	}
	// 0 表示保存到数据库中
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := system.SysResource{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, i.Upload(f)
	}
	return
}

func (i *ResourceController) Upload(file system.SysResource) error {
	return global.MEIS_DB.Create(&file).Error
}

// 获取资源列表
func (i *ResourceController) GetResourceList(info request.ListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MEIS_DB.Model(&system.SysResource{})
	var resourceList []system.SysResource
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&resourceList).Error
	return resourceList, total, err
}

// 删除文件
func (i *ResourceController) RemoveFile(file system.SysResource) (err error) {
	var fileFromDb system.SysResource
	fileFromDb, err = i.FindFile(file.ID)
	if err != nil {
		return
	}
	oss := upload.NewOOS()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	// 增加这个属性{Unscoped}就是强删除
	err = global.MEIS_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// 编辑文件
func (i *ResourceController) UpdateFile(file system.SysResource) (err error) {
	var fileFromDb system.SysResource
	return global.MEIS_DB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

// 查找文件
func (i *ResourceController) FindFile(id uint) (system.SysResource, error) {
	var file system.SysResource
	err := global.MEIS_DB.Where("id = ?", id).First(&file).Error
	return file, err
}
