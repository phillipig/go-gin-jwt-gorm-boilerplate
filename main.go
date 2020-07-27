package main

import (
	"fmt"
	"go-api/Config"
	"go-api/Database"
	"go-api/Models"
	"go-api/Routes"
	"time"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	//Gorm
	Database.Mysql, err = gorm.Open("mysql", Database.DbURL(Database.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Database.Mysql.Close()
	Database.Mysql.AutoMigrate(&Models.User{})

	//Pool
	Database.Mysql.DB().SetMaxIdleConns(1)
	Database.Mysql.DB().SetMaxOpenConns(10)
	Database.Mysql.DB().SetConnMaxLifetime(time.Hour)

	//Gin
	r := Routes.SetupRouter()
	r.Run(":" + Config.GetEnvKey("API_PORT"))
}
