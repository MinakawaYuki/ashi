package service

import (
	"ashi/model"
	"ashi/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

var log utils.Log

type Result struct {
	model.Base
	Name     string `json:"name,omitempty"`
	UserName string `json:"username,omitempty"`
}

func AddUser(c *gin.Context) error {
	var user model.User
	params, _ := c.Get("params")
	param := utils.AnyMapToStringMap(params.(map[string]interface{}))
	user.Name = param["name"]
	user.UserName = param["username"]
	user.PassWord = param["password"]
	salt := utils.GenRandString(6)
	user.PassWord = utils.Md5(user.PassWord + salt)
	user.Salt = salt

	err := user.Add(user)
	if err != nil {
		log.Info(map[string]interface{}{}, "user 创建失败")
		return err
	}
	return nil
}

func Login(username string, password string) (data map[string]interface{}, err error) {
	var userModel model.User
	user, err := userModel.GetByPw(username)
	if err != nil {
		log.Info(map[string]interface{}{
			"username": username,
		}, "查询失败")
		return map[string]interface{}{}, err
	}
	passwordMd5 := utils.Md5(password + user.Salt)

	if passwordMd5 != user.PassWord {
		newErr := errors.New("密码不正确")
		return map[string]interface{}{}, newErr
	}

	// 获取token
	data = make(map[string]interface{})
	token, err := utils.GenToken(utils.Auth{
		Username: user.UserName,
		Password: user.PassWord,
	})

	if err != nil {
		log.Info(map[string]interface{}{
			"username": username,
		}, "查询失败")
		return map[string]interface{}{}, err
	}
	var result Result
	result.Base.ID = user.ID
	result.Base.CreatedAt = user.CreatedAt
	result.Base.UpdatedAt = user.UpdatedAt
	result.Name = user.Name
	result.UserName = user.UserName
	data["user"] = result
	data["token"] = token

	return data, nil
}
