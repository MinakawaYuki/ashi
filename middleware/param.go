package middleware

import (
	"ashi/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取query 参数
		var params map[string]any
		query := c.Request.URL.Query()
		var queryMap = make(map[string]any, len(query))
		for k := range query {
			queryMap[k] = c.Query(k)
		}
		// 获取post 参数
		var postMap = make(map[string]any, len(c.Request.PostForm))
		for k, v := range c.Request.PostForm {
			if len(v) > 1 {
				postMap[k] = v
			} else if len(v) == 1 {
				postMap[k] = v[0]
			}
		}
		// 获取json内容
		json := make(map[string]any)
		c.BindJSON(&json)
		jsonMap := make(map[string]any)
		for k1, v2 := range json {
			jsonMap[k1] = v2
		}
		params = utils.MergeMap(queryMap, postMap, jsonMap)
		fmt.Println("params:", params)
		c.Set("params", params)
		c.Next()
	}
}
