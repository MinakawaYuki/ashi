package model

type Summon struct {
	Base
	Name     string      `json:"name" gorm:"name"`
	SummonID string      `json:"summon_id" gorm:"summon_id"`
	Property string      `json:"property" gorm:"property"`
	Thumb    string      `json:"thumb" gorm:"thumb"`
	Pic      []SummonPic `json:"pic" gorm:"foreignKey:SummonID;references:SummonID"`
}

func (Summon) TableName() string {
	return "summon"
}
