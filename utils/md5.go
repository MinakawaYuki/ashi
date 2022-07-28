package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

var log Log

func Md5(str string) string {
	m := md5.New()
	_, err := m.Write([]byte(str))
	if err != nil {
		log.Err(map[string]interface{}{}, "md5 err :"+err.Error())
		os.Exit(1)
	}
	md5String := hex.EncodeToString(m.Sum(nil))
	return md5String
}
