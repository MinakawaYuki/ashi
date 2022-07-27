package model

import "ashi/setting"

type Character struct {
	Base
	Name              string         `json:"name" gorm:"column:name"`
	CharacterID       string         `json:"character_id" gorm:"column:character_id"`
	CharacterCategory string         `json:"character_category" gorm:"column:character_category"`
	CharacterType     string         `json:"character_type" gorm:"column:character_type"`
	CharacterGender   string         `json:"character_gender" gorm:"column:character_gender"`
	CharacterWeapon   string         `json:"character_weapon" gorm:"column:character_weapon"`
	CharacterRace     string         `json:"character_race" gorm:"column:character_race"`
	CharacterThumb    string         `json:"character_thumb" gorm:"column:character_thumb"`
	Property          string         `json:"property" gorm:"column:property"`
	Pic               []CharacterPic `json:"pic" gorm:"foreignKey:CharacterID;references:CharacterID"`
}

func (Character) TableName() string {
	return "character"
}

func (c *Character) GetCharacterByID(id string) (character Character, err error) {
	err = setting.Db.Model(&Character{}).Where("character_id = ?", id).Preload("Pic").First(&character).Error
	return character, err
}

func (c *Character) GetCharacter(query map[string]interface{}) (character Character, err error) {
	err = setting.Db.Where(query).Preload("Pic").Find(&character).Error
	return character, err
}

func (c *Character) GetCharacterList(query map[string]interface{}, page int, pageSize int) (characters []Character, err error) {
	err = setting.Db.Where(query).Offset((page - 1) * pageSize).Limit(pageSize).Preload("Pic").Find(&characters).Error
	return characters, err
}

func (c *Character) GetCount(query map[string]interface{}) (count int64) {
	setting.Db.Model(&Character{}).Where(query).Count(&count)
	return count
}
