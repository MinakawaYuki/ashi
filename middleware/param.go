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
		// 获取post 参数  Content-Type 需为 multipart/form-data
		var postMap = make(map[string]any, len(c.Request.PostForm))
		if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
			fmt.Println("获取post 参数错误：", err.Error())
		} else {
			for k, v := range c.Request.PostForm {
				if len(v) > 1 {
					postMap[k] = v
				} else if len(v) == 1 {
					postMap[k] = v[0]
				}
			}
		}
		params = utils.MergeMap(queryMap, postMap)
		fmt.Println("params:", params)
		c.Set("params", params)
		c.Next()
	}
}
