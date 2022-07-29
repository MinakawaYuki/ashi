package model

type WeaponPic struct {
	Base
	WeaponID string `json:"weapon_id" gorm:"weapon_id"`
	Name     string `json:"name" gorm:"name"`
	Thumb    string `json:"thumb" gorm:"thumb"`
}

func (WeaponPic) TableName() string {
	return "weapon_pic"
}
