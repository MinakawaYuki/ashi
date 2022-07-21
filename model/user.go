package model

import (
	"ashi/setting"
)

type User struct {
	Base
	Name     string `json:"name"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func (u User) Add(model interface{}) error {
	err := setting.Db.DB.Create(model).Error
	return err
}
