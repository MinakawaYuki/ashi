package model

type Job struct {
	Base
	Name        string `json:"name" gorm:"column:name"`
	Icon        string `json:"icon" gorm:"column:icon"`
	ThumbGran   string `json:"thumb_gran" gorm:"column:thumb_gran"`
	ThumbDjeeta string `json:"thumb_djeeta" gorm:"column:thumb_djeeta"`
}

func (Job) TableName() string {
	return "job"
}
