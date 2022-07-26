package model

import (
	"ashi/setting"
)

type User struct {
	Base
	Name     string `gorm:"column:name" json:"name,omitempty"`
	UserName string `gorm:"column:username" json:"username,omitempty" binding:"required"`
	PassWord string `gorm:"column:password" json:"password,omitempty" binding:"required"`
	Salt     string `gorm:"column:salt" json:"salt,omitempty"`
}

func (u User) Add(model User) error {
	err := setting.Db.Create(&model).Error
	return err
}

func (u User) GetByPw(username string) (user User, err error) {
	err = setting.Db.Model(&User{}).Where("username = ?", username).Find(&user).Error
	return user, err
}
