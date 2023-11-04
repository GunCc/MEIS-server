package system

import (
	"MEIS-server/global"
	sysReq "MEIS-server/model/system/request"

	"MEIS-server/model/system"
	"MEIS-server/utils/upload"
	"errors"
	"mime/multipart"
	"strings"

	"gorm.io/gorm"
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
		err = i.Upload(&f)
		return f, err
	}
	return
}

func (i *ResourceController) Upload(file *system.SysResource) error {
	return global.MEIS_DB.Create(file).Error
}

// 获取资源列表
func (i *ResourceController) GetResourceList(info sysReq.SysFileListInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword

	db := global.MEIS_DB.Model(&system.SysResource{})
	var resourceList []system.SysResource

	db = db.Limit(limit).Offset(offset)
	type_id := info.TypeId

	if !errors.Is(global.MEIS_DB.Where("id = ? ", type_id).First(&system.SysResourceType{}).Error, gorm.ErrRecordNotFound) {
		db = db.Where("sys_resource_type_id = ?", type_id)
	}
	db = db.Where("name LIKE ?", "%"+keyword+"%")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&resourceList).Error

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

// 获取文件分类
func (i *ResourceController) GetResourceTypeList() (list interface{}, total int64, err error) {
	db := global.MEIS_DB.Model(&system.SysResourceType{})
	var resourceList []system.SysResourceType
	err = db.Count(&total).Find(&resourceList).Error
	return resourceList, total, err
}

// 添加文件分类
func (i *ResourceController) AddFileType(fileType *system.SysResourceType) (err error) {
	if !errors.Is(global.MEIS_DB.Where("name = ?", fileType.Name).First(&system.SysResourceType{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("分类名重复")
	}
	err = global.MEIS_DB.Create(fileType).Error

	return err
}

// 编辑文件分类
func (i *ResourceController) UpdateFileType(fileType system.SysResourceType) (err error) {
	var fileTypeFromDb system.SysResourceType
	return global.MEIS_DB.Where("id = ?", fileType.ID).First(&fileTypeFromDb).Update("name", fileType.Name).Error
}

// 删除文件分类
func (i *ResourceController) DeleteFileType(fileTypeId uint) (err error) {
	var fileTypeFromDb system.SysResourceType
	err = global.MEIS_DB.Where("id = ?", fileTypeId).Unscoped().Delete(&fileTypeFromDb).Error

	return err
}

// 文件绑定类型
func (i *ResourceController) FileBindType(info sysReq.SysFileBindType) (err error) {
	db := global.MEIS_DB.Model(system.SysResource{})

	err = db.Where("id in ?", info.SysResourceId).Updates(system.SysResource{SysResourceTypeId: info.TypeId}).Error
	return err
}
