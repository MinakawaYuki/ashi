package model

import "ashi/setting"

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

func (w *Weapon) GetList(j Weapon, page int, pageSize int) (weapon []Weapon, err error) {
	err = setting.Db.Where(&j).Offset((page - 1) * pageSize).Limit(pageSize).Preload("Pic").Find(&weapon).Error
	return weapon, err
}

func (w *Weapon) GetCount(j Weapon) (count int64) {
	setting.Db.Model(&Weapon{}).Where(&j).Count(&count)
	return count
}
