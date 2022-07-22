package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	Mysql       Mysql
	Server      Server
	RedisStruct RedisStruct
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	DataBase string `json:"dataBase"`
	TimeOut  string `json:"timeOut"`
}

type Server struct {
	Port        string `json:"port"`
	ReleaseMode string `json:"releaseMode"`
}

type RedisStruct struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

func (c *Config) GetConfig() Config {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		os.Exit(1)
	}
	redisDb, err := cfg.Section("redis").Key("DB").Int()
	config := Config{
		Mysql: Mysql{
			Host:     cfg.Section("mysql").Key("Host").String(),
			Port:     cfg.Section("mysql").Key("Port").String(),
			UserName: cfg.Section("mysql").Key("UserName").String(),
			PassWord: cfg.Section("mysql").Key("PassWord").String(),
			DataBase: cfg.Section("mysql").Key("DataBase").String(),
			TimeOut:  cfg.Section("mysql").Key("TimeOut").String(),
		},
		Server: Server{
			Port:        cfg.Section("server").Key("Port").String(),
			ReleaseMode: cfg.Section("server").Key("ReleaseMode").String(),
		},
		RedisStruct: RedisStruct{
			Host:     cfg.Section("redis").Key("Host").String(),
			Port:     cfg.Section("redis").Key("Port").String(),
			Password: cfg.Section("redis").Key("Password").String(),
			DB:       redisDb,
		},
	}

	return config
}
