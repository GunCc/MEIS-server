package oa

import (
	"MEIS-server/global"
)

// 员工
type OATrain struct {
	global.MEIS_MODEL
	TrainName   string      `json:"train_name" gorm:"comment:培训名称"`
	TrainDesc   string      `json:"train_desc" gorm:"comment:培训说明"`
	IsApart     int         `json:"is_apart" gorm:"comment:是否参加"`
	Reason      string      `json:"reason" gorm:"comment:没有参加的原因"`
	PersonnelID uint        `json:"personnel_id" gorm:"comment:员工ID"`
	OAPersonnel OAPersonnel `json:"personnel" gorm:"foreignKey:ID;comment:员工"`
}

func (OATrain) TableName() string {
	return "oa_train"
}
