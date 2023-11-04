package request

type SysFileBindType struct {
	SysResourceId []uint `json:"files"`
	TypeId        uint   `json:"type_id" `
}
