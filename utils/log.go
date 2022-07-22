package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Log struct {
	logrus.Logger
}

func (l *Log) Init() {
	l.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)
}

func (l *Log) Warning(field map[string]interface{}, msg string) {
	file, err := os.OpenFile("./runtime/logs/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("log warning err :", err.Error())
	}
	l.Out = file
	l.WithFields(field).Warn(msg)
}
