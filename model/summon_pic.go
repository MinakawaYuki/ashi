package model

type SummonPic struct {
	Base
	SummonID string `json:"summon_id" gorm:"summon_id"`
	Name     string `json:"name" gorm:"name"`
	Thumb    string `json:"thumb" gorm:"thumb"`
}

func (SummonPic) TableName() string {
	return "summon_pic"
}
