package model

type Base struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" form:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
