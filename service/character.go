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
	fmt.Println("ch ", ch)
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
		err = json.Unmarshal([]byte(url.Thumb), &thumb)
		fmt.Println("thumb ", thumb)
		if err != nil {
			return CharacterRes{}, errors.New("json 解析 Thumb 出错")
		}
		err = json.Unmarshal([]byte(url.Skin), &skin)
		if err != nil {
			return CharacterRes{}, errors.New("json 解析 Skin 出错")
		}
		for _, t := range thumb {
			res.Pic = append(res.Pic, t)
		}
		for _, s := range skin {
			res.Pic = append(res.Pic, s)
		}
	}
	return res, nil
}
