package oa

import (
	"MEIS-server/global"
	"time"
)

// 员工
type OATrain struct {
	global.MEIS_MODEL
	TrainName string    `json:"train_name" gorm:"comment:培训名称"`
	TrainDesc string    `json:"train_desc" gorm:"comment:培训说明"`
	StartTime time.Time `json:"start_time" gorm:"comment:预计开始时间"`
	EndTime   time.Time `json:"end_time" gorm:"comment:预计完成时间"`
	IsApart   bool      `json:"is_apart" gorm:"comment:是否参加"`
	Reason    string    `json:"reason" gorm:"comment:没有参加的原因"`
}

func (OATrain) TableName() string {
	return "oa_train"
}
