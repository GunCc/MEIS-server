package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system"
	"MEIS-server/utils/upload"
	"mime/multipart"
	"strings"
)

type ResourceController struct {
}

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
