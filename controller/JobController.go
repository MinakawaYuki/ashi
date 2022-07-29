package controller

import (
	"ashi/service"
	"github.com/gin-gonic/gin"
)

type JobController struct {
}

var jj service.JobRes

func (j *JobController) GetList(c *gin.Context) {
	jr, total, err := jj.GetList(c)
	if err != nil {
		ApiError("查询失败", map[string]interface{}{
			"err": err.Error(),
		}, c)
		return
	}
	ApiSuccess("成功", map[string]interface{}{
		"list":  jr,
		"total": total,
	}, c)
}
