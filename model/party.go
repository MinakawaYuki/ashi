package model

type Party struct {
	Base
	PartyNo        string `json:"party_no" gorm:"column:party_no"`                 // 队伍编号
	ActivityId     int    `json:"activity_id" gorm:"column:activity_id"`           // 适用活动id 为空/0 时代表日常
	CharacterJson  string `json:"character_json" gorm:"column:character_json"`     // 队伍角色构成json
	WeaponJson     string `json:"weapon_json" gorm:"column:weapon_json"`           // 武器盘json
	SummonJson     string `json:"summon_json" gorm:"column:summon_json"`           // 召唤石json
	JobId          int    `json:"job_id" gorm:"column:job_id"`                     // 主角职业
	JobAbilityJson string `json:"job_ability_json" gorm:"column:job_ability_json"` // 主角携带技能
	Remark         string `json:"remark" gorm:"column:remark"`                     // 队伍备注
}

func (Party) TableName() string {
	return "party"
}
