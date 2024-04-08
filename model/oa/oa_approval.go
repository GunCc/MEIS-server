package oa

import (
	"MEIS-server/global"
)

const (
	APPROVAL_SALARY     = 5  // 薪资审核
	APPROVAL_ATTENDANCE = 10 // 考勤审核
)

// 员工
type OAApproval struct {
	global.MEIS_MODEL
	ApprovalTypeID int `json:"approval_type_id" gorm:"comment:审批类型的ID"`
	ApprovalType   int `json:"approval_type" gorm:"comment:审批类型"`
	IsPast         int `json:"is_past" gorm:"是否审批通过，未审核0 通过1 未通过2"`
}

func (OAApproval) TableName() string {
	return "oa_approvals"
}
