package setting

import (
	"fmt"
	redisClient "github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var Db *gorm.DB
var Redis *redisClient.Client

func init() {
	InitMysql()
	InitRedis()
}

func InitMysql() {
	var c Config
	config := c.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Mysql.UserName, config.Mysql.PassWord, config.Mysql.Host, config.Mysql.Port, config.Mysql.DataBase)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		fmt.Println("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
		os.Exit(1)
	}

	_, err = conn.DB()
	if err != nil {
		fmt.Println("[SetupDefaultDatabase#sqlDb error]: " + err.Error() + " " + dsn)
		os.Exit(1)
	}
	fmt.Printf("\nDefault database connection successful: %s\n", dsn)

	Db = conn
}

func InitRedis() {
	var c Config
	config := c.GetConfig()
	rdb := redisClient.NewClient(
		&redisClient.Options{Addr: config.RedisStruct.Host + ":" + config.RedisStruct.Port, Password: config.RedisStruct.Password, DB: config.RedisStruct.DB})
	fmt.Printf("\nDefault redis connection successful\n")
	Redis = rdb
}
