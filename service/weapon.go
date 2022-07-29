package service

import (
	"ashi/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type WeaponRes struct {
	model.Base
	Name     string   `json:"name"`
	NameJp   string   `json:"name_jp"`
	WeaponID string   `json:"weapon_id"`
	Property string   `json:"property"`
	Skills   []string `json:"skills"`
	Thumb    []string `json:"thumb"`
}

func (WeaponRes) GetList(c *gin.Context) (res []WeaponRes, total int64, err error) {
	var w model.Weapon
	var wr []model.Weapon
	var ress []WeaponRes
	err = c.ShouldBindJSON(&w)
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "15")).MustInt()
	if err != nil {
		return []WeaponRes{}, 0, err
	}
	wr, err = w.GetList(w, page, pageSize)
	if err != nil {
		return []WeaponRes{}, 0, err
	}
	for _, item := range wr {
		var res WeaponRes
		res = res.makeRes(item)
		ress = append(ress, res)
	}
	return ress, w.GetCount(w), err
}

func (WeaponRes) makeRes(w model.Weapon) WeaponRes {
	var res WeaponRes
	res.Name = w.Name
	res.NameJp = w.NameJp
	res.WeaponID = w.WeaponID
	res.Property = w.Property
	var sr []string
	err := json.Unmarshal([]byte(w.Skills), &sr)
	if err != nil {
		log.Info(map[string]interface{}{}, "json 解析 skill 出错")
	}
	for _, r := range sr {
		res.Skills = append(res.Skills, r)
	}

	for _, url := range w.Pic {
		var thumb []string
		err = json.Unmarshal([]byte(url.Thumb), &thumb)
		if err != nil {
			log.Info(map[string]interface{}{}, "json 解析 Thumb 出错")
			continue
		}
		for _, t := range thumb {
			res.Thumb = append(res.Thumb, t)
		}
	}
	return res
}
