package databases

import (
	"fmt"
	"go-api/config"
	"go-api/models"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

var once sync.Once
var mysql *gorm.DB
var env config.EnvMap
var err error

func NewMysql() *gorm.DB {
	once.Do(func() {
		env = config.NewEnv()
		mysql, err = gorm.Open(env.GetEnvKey("API_GORM_DIALECT"), urlMysql())
		if err == nil {
			configMysql()
			migrationsMysql()
		} else {
			log.Fatal("GORM: ", err)
		}
	})
	return mysql
}

func configMysql() {
	mysql.LogMode(env.GetEnvKeyBool("API_GORM_LOG"))
	mysql.SingularTable(true)
	mysql.DB().SetMaxIdleConns(2)
	mysql.DB().SetMaxOpenConns(10)
	mysql.DB().SetConnMaxLifetime(time.Hour)
}

func migrationsMysql() {
	mysql.AutoMigrate(&models.User{})
}

func urlMysql() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		env.GetEnvKey("API_GORM_USER"),
		env.GetEnvKey("API_GORM_PASS"),
		env.GetEnvKey("API_GORM_HOST"),
		env.GetEnvKeyInt("API_GORM_PORT"),
		env.GetEnvKey("API_GORM_DB"),
	)
}
