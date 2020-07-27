package Databases

import (
	"fmt"
	"go-api/Config"
	"go-api/Models"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

var Mysql *gorm.DB
var err error

func InitMysql() {
	Mysql, err = gorm.Open(Config.GetEnvKey("API_GORM_DIALECT"), urlMysql())
	if err == nil {
		configMysql()
		migrationsMysql()
	} else {
		log.Fatal("GORM: ", err)
	}
}

func configMysql() {
	Mysql.LogMode(Config.GetEnvKeyBool("API_GORM_LOG"))
	Mysql.SingularTable(true)
	Mysql.DB().SetMaxIdleConns(2)
	Mysql.DB().SetMaxOpenConns(10)
	Mysql.DB().SetConnMaxLifetime(time.Hour)
}

func migrationsMysql() {
	Mysql.AutoMigrate(&Models.User{})
}

func urlMysql() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Config.GetEnvKey("API_GORM_USER"),
		Config.GetEnvKey("API_GORM_PASS"),
		Config.GetEnvKey("API_GORM_HOST"),
		Config.GetEnvKeyInt("API_GORM_PORT"),
		Config.GetEnvKey("API_GORM_DB"),
	)
}
