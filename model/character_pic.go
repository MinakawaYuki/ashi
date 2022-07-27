package model

type CharacterPic struct {
	Base
	CharacterID string `json:"character_id" gorm:"character_id"`
	Name        string `json:"name" gorm:"name"`
	Thumb       string `json:"thumb" gorm:"thumb"`
	Skin        string `json:"skin" gorm:"skin"`
}

func (CharacterPic) TableName() string {
	return "character_pic"
}
