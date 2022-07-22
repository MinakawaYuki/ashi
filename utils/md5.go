package utils

import (
	"ashi/setting"
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func Test() {
	fmt.Println(setting.Redis)
}
