package controller

import (
	"ashi/service"
	"github.com/gin-gonic/gin"
)

type CharacterController struct {
}

var chh service.CharacterRes

func (ch *CharacterController) GetCharacterById(c *gin.Context) {
	var chr service.CharacterRes
	chr, err := chh.GetByID(c)
	if err != nil {
		ApiError("查询失败", map[string]interface{}{
			"err": err.Error(),
		}, c)
		return
	}
	ApiSuccess("成功", map[string]interface{}{
		"data": chr,
	}, c)
}

func (ch *CharacterController) GetCharacter(c *gin.Context) {
	var chr service.CharacterRes
	chr, err := chh.GetDetail(c)
	if err != nil {
		ApiError("查询失败", map[string]interface{}{
			"err": err.Error(),
		}, c)
		return
	}
	ApiSuccess("成功", map[string]interface{}{
		"data": chr,
	}, c)
}

func (ch *CharacterController) GetCharacterList(c *gin.Context) {
	var chr []service.CharacterRes
	chr, total, err := chh.GetList(c)
	if err != nil {
		ApiError("查询失败", map[string]interface{}{
			"err": err.Error(),
		}, c)
		return
	}
	ApiSuccess("成功", map[string]interface{}{
		"list":  chr,
		"total": total,
	}, c)
}
