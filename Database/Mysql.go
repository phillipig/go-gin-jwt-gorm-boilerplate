package Database

import (
	"fmt"
	"go-api/Config"

	"github.com/jinzhu/gorm"
)

var Mysql *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     Config.GetEnvKey("API_MYSQL_HOST"),
		Port:     Config.GetEnvKeyInt("API_MYSQL_PORT"),
		User:     Config.GetEnvKey("API_MYSQL_USER"),
		Password: Config.GetEnvKey("API_MYSQL_PASS"),
		DBName:   Config.GetEnvKey("API_MYSQL_DB"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
