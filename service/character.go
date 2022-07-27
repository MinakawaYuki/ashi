package service

import (
	"ashi/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CharacterRes struct {
	model.Base
	Name              string   `json:"name"`
	CharacterID       string   `json:"character_id"`
	CharacterCategory string   `json:"character_category"`
	CharacterType     string   `json:"character_type"`
	CharacterGender   string   `json:"character_gender"`
	CharacterWeapon   string   `json:"character_weapon"`
	CharacterRace     string   `json:"character_race"`
	CharacterThumb    string   `json:"character_thumb"`
	Property          string   `json:"property"`
	Pic               []string `json:"pic"`
}
type query struct {
	Character model.Character
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
}

func GetByID(c *gin.Context) (res CharacterRes, err error) {
	id, exist := c.GetQuery("character_id")
	if !exist {
		return CharacterRes{}, errors.New("请传入id")
	}
	var ch model.Character
	ch, err = ch.GetCharacterByID(id)
	if err != nil {
		return CharacterRes{}, err
	}

	res = res.makeRes(ch)
	return res, nil
}

func GetDetail(c *gin.Context) (res CharacterRes, err error) {
	var ch model.Character
	err = c.ShouldBindJSON(&ch)
	if err != nil {
		return CharacterRes{}, err
	}
	ch, err = ch.GetCharacter(ch)
	fmt.Println("ch res :", ch)
	if err != nil {
		return CharacterRes{}, err
	}
	res = res.makeRes(ch)
	return res, err
}

func GetList(c *gin.Context) (res []CharacterRes, total int64, err error) {
	var ch query
	var character model.Character
	var chr []model.Character
	var ress []CharacterRes
	err = c.ShouldBindJSON(&ch)
	if err != nil {
		return []CharacterRes{}, 0, err
	}
	chr, err = character.GetCharacterList(ch.Character, ch.Page, ch.PageSize)
	if err != nil {
		return []CharacterRes{}, 0, err
	}
	for _, item := range chr {
		var res CharacterRes
		res = res.makeRes(item)
		ress = append(ress, res)
	}
	return ress, character.GetCount(ch.Character), err
}

func (CharacterRes) makeRes(ch model.Character) CharacterRes {
	var res CharacterRes
	res.CharacterID = ch.CharacterID
	res.Name = ch.Name
	res.CharacterCategory = ch.CharacterCategory
	res.CharacterType = ch.CharacterType
	res.CharacterGender = ch.CharacterGender
	res.CharacterWeapon = ch.CharacterWeapon
	res.CharacterRace = ch.CharacterRace
	res.CharacterThumb = ch.CharacterThumb
	res.Property = ch.Property

	for _, url := range ch.Pic {
		var thumb []string
		var skin []string
		err := json.Unmarshal([]byte(url.Thumb), &thumb)
		fmt.Println("thumb ", thumb)
		if err != nil {
			log.Info(map[string]interface{}{}, "json 解析 Thumb 出错")
			continue
		}
		err = json.Unmarshal([]byte(url.Skin), &skin)
		if err != nil {
			log.Info(map[string]interface{}{}, "json 解析 Skin 出错")
			continue
		}
		for _, t := range thumb {
			res.Pic = append(res.Pic, t)
		}
		for _, s := range skin {
			res.Pic = append(res.Pic, s)
		}
	}

	return res
}
