package utils

import (
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"time"
)

const fileRouter = "/upload/"

func init() {
	_, err := os.Stat("./runtime/upload")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll("./runtime/upload", 0666)
		}
	}
}

func UploadFile(c *gin.Context) (filePath string, err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	extName := path.Ext(file.Filename)
	filename := getFileName(file.Filename) + extName
	dst := path.Join("./runtime/upload", filename)
	saveErr := c.SaveUploadedFile(file, dst)
	if saveErr != nil {
		return "", saveErr
	}
	return fileRouter + filename, nil
}

func UploadFiles(c *gin.Context) (filePath []string, err error) {
	var urls []string
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		extName := path.Ext(file.Filename)
		filename := getFileName(file.Filename) + extName
		dst := path.Join("./runtime/upload", filename)
		urls = append(urls, fileRouter+filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return make([]string, 0), err
		}
	}
	return urls, nil
}

func getFileName(filename string) string {
	//创建图片保存目录,linux下需要设置权限（0666可读可写） runtime/upload/xxx
	day := time.Now().Format("20060102")
	name := Md5(day + filename)
	return name
}
