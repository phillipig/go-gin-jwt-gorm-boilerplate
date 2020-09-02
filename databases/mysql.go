package databases

import (
	"fmt"
	"go-api/configs"
	"go-api/models"
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once
var db *gorm.DB
var env configs.EnvMap
var err error

func NewMysql() *gorm.DB {
	once.Do(func() {
		env = configs.NewEnv()
		db, err = connect()
		if err == nil {
			config()
			migrations()
		} else {
			log.Fatal("GORM: ", err)
		}
	})
	return db
}

func connect() (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       url(), // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
}

func url() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		env.GetEnvKey("API_GORM_USER"),
		env.GetEnvKey("API_GORM_PASS"),
		env.GetEnvKey("API_GORM_HOST"),
		env.GetEnvKeyInt("API_GORM_PORT"),
		env.GetEnvKey("API_GORM_DB"),
	)
}

func config() {
	sqlDB, _ := db.DB()                   // Get generic database object sql.DB to use its functions
	sqlDB.SetMaxIdleConns(1)              // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(10)             // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxIdleTime(time.Minute) // SetConnMaxIdleTime sets the maximum idle time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)   // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
}

func migrations() {
	db.AutoMigrate(&models.User{})
}
