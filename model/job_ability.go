package model

type JobAbility struct {
	Base
	JobId  int    `json:"job_id" gorm:"column:job_id"`
	Name   string `json:"name" gorm:"column:name"`
	NameCn string `json:"name_cn" gorm:"column:name_cn"` // 中文技能名
	Icon   string `json:"icon" gorm:"column:icon"`       // 图标
}

func (JobAbility) TableName() string {
	return "job_ability"
}
