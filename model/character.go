package model

type Character struct {
	Base
	Name              string `json:"name" gorm:"column:name"`
	CharacterId       string `json:"character_id" gorm:"column:character_id"`
	CharacterCategory string `json:"character_category"`
	CharacterType     string `json:"character_type"`
	CharacterGender   string `json:"character_gender"`
	CharacterWeapon   string `json:"character_weapon"`
	CharacterRace     string `json:"character_race"`
	CharacterThumb    string `json:"character_thumb"`
	Property          string `json:"property"`
}
