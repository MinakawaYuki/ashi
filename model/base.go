package model

import "time"

type Base struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp" json:"updated_at,omitempty"`
}
