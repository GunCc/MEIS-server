package upload

import (
	"MEIS-server/global"
	"mime/multipart"
)

// 对象存储接口 {外接的}
type OOS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// 根据配置文件实例化一个OOS接口
func NewOOS() OOS {
	switch global.MEIS_CONFIG.System.OOSType {
	case "local":
		return &Local{}
	default:
		return &Local{}
	}
}
