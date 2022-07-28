package model

type Activity struct {
	Base
	Name   string `json:"name" gorm:"column:name"`     // 活动名称
	Banner string `json:"banner" gorm:"column:banner"` // banner
}

func (Activity) TableName() string {
	return "activity"
}
