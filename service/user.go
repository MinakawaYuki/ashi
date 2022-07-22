package service

import (
	"ashi/model"
	"ashi/utils"
	"github.com/gin-gonic/gin"
)

var log utils.Log

func AddUser(c *gin.Context) error {
	var user model.User
	err := c.ShouldBindJSON(&user)
	salt := utils.GenRandString(6)
	user.PassWord = utils.Md5(user.PassWord + salt)
	user.Salt = salt
	if err != nil {
		log.Info(map[string]interface{}{}, "user 绑定json失败")
		return err
	}
	err = user.Add(user)
	if err != nil {
		log.Info(map[string]interface{}{}, "user 创建失败")
		return err
	}
	return nil
}
