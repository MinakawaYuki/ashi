package controller

import (
	"ashi/service"
	"github.com/gin-gonic/gin"
)

type WeaponController struct {
}

var ws service.WeaponRes

func (w *WeaponController) GetList(c *gin.Context) {
	wr, total, err := ws.GetList(c)
	if err != nil {
		ApiError("查询失败", map[string]interface{}{
			"err": err.Error(),
		}, c)
		return
	}
	ApiSuccess("成功", map[string]interface{}{
		"list":  wr,
		"total": total,
	}, c)
}
