package setting

import (
	"ashi/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

type mysqlUtil struct {
	DB *gorm.DB
}

var Db mysqlUtil

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
	fmt.Printf("\nDefault atabase connection successful: %s\n", dsn)

	Db = mysqlUtil{DB: conn}
}

func InitServer() {
	var c Config
	config := c.GetConfig()
	r := gin.Default()
	gin.SetMode(config.Server.ReleaseMode)
	router.Register(r)
	r.Run(":" + config.Server.Port)
}
