package model

import (
	"ashi/setting"
)

type User struct {
	Base
	Name     string `gorm:"column:name"`
	UserName string `gorm:"column:username"`
	PassWord string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
}

func (u User) Add(model User) error {
	err := setting.Db.Create(&model).Error
	return err
}
