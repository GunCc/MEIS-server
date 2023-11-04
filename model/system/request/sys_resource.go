package request

import "MEIS-server/model/commen/request"

type SysFileBindType struct {
	SysResourceId []uint `json:"files"`
	TypeId        uint   `json:"type_id" `
}

type SysFileListInfo struct {
	request.ListInfo
	TypeId string `json:"type_id"`
}
