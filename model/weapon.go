package model

type Weapon struct {
	Base
	Name     string      `json:"name" gorm:"name"`
	NameJp   string      `json:"name_jp" gorm:"name_jp"`
	WeaponID string      `json:"weapon_id" gorm:"weapon_id"`
	Property string      `json:"property" gorm:"property"`
	Skills   string      `json:"skills" gorm:"skills"`
	Pic      []WeaponPic `json:"pic" gorm:"foreignKey:WeaponID;references:WeaponID"`
}

func (Weapon) TableName() string {
	return "weapon"
}
