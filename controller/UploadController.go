package controller

import (
	"ashi/utils"
	"github.com/gin-gonic/gin"
)

type UploadController struct {
}

func (u *UploadController) UploadFile(c *gin.Context) {
	path, err := utils.UploadFile(c)
	if err != nil {
		ApiError("上传失败", map[string]interface{}{
			"msg": err.Error(),
		}, c)
		return
	}
	ApiSuccess("上传成功", map[string]interface{}{
		"path": path,
	}, c)
}

func (u *UploadController) UploadFiles(c *gin.Context) {
	path, err := utils.UploadFiles(c)
	if err != nil {
		ApiError("上传失败", map[string]interface{}{
			"msg": err.Error(),
		}, c)
		return
	}
	ApiSuccess("上传成功", map[string]interface{}{
		"data": path,
	}, c)
}
