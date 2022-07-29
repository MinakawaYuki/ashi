package service

import (
	"ashi/model"
	"github.com/gin-gonic/gin"
)

type JobRes struct {
	model.Base
	Name        string                   `json:"name"`
	Icon        string                   `json:"icon"`
	ThumbGran   string                   `json:"thumb_gran"`
	ThumbDjeeta string                   `json:"thumb_djeeta"`
	Ability     []map[string]interface{} `json:"ability"`
}

func (JobRes) GetList(c *gin.Context) (res []JobRes, total int64, err error) {
	var j model.Job
	var jr []model.Job
	var ress []JobRes
	err = c.ShouldBindJSON(&j)
	if err != nil {
		return []JobRes{}, 0, err
	}
	jr, err = j.GetJobList(j)
	if err != nil {
		return []JobRes{}, 0, err
	}
	for _, item := range jr {
		var res JobRes
		res = res.makeRes(item)
		ress = append(ress, res)
	}
	return ress, j.GetCount(j), err
}

func (JobRes) makeRes(ch model.Job) JobRes {
	var res JobRes
	res.Name = ch.Name
	res.Icon = ch.Icon
	res.ThumbGran = ch.ThumbGran
	res.ThumbDjeeta = ch.ThumbDjeeta

	for _, url := range ch.Ability {
		var skill = make(map[string]interface{})
		skill["name"] = url.Name
		skill["name_cn"] = url.NameCn
		skill["type"] = url.Type
		skill["icon"] = url.Icon

		res.Ability = append(res.Ability, skill)

	}
	return res
}
