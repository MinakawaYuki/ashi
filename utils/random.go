package utils

import (
	"math/rand"
	"time"
)

var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenRandString(n int) string {
	rand.Seed(time.Now().Unix())

	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		log.Err(map[string]interface{}{}, "生成随机字符串失败:"+err.Error())
		return ""
	}
	for i, x := range b {
		arc = x & 61
		b[i] = longLetters[arc]
	}
	return string(b)
}
