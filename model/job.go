package model

import "ashi/setting"

type Job struct {
	Base
	Name        string       `json:"name" gorm:"column:name"`
	Icon        string       `json:"icon" gorm:"column:icon"`
	ThumbGran   string       `json:"thumb_gran" gorm:"column:thumb_gran"`
	ThumbDjeeta string       `json:"thumb_djeeta" gorm:"column:thumb_djeeta"`
	Ability     []JobAbility `json:"ability" gorm:"foreignKey:JobID"`
}

func (Job) TableName() string {
	return "job"
}

func (c *Job) GetJobList(j Job) (jobs []Job, err error) {
	err = setting.Db.Where(&j).Preload("Ability").Find(&jobs).Error
	return jobs, err
}

func (c *Job) GetCount(j Job) (count int64) {
	setting.Db.Model(&Job{}).Where(&j).Count(&count)
	return count
}
