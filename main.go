package main

import (
	"go-api/Config"
	"go-api/Databases"
	"go-api/Routes"
)

func main() {
	//Env
	Config.InitEnv()

	//Gorm
	Databases.InitMysql()
	defer Databases.Mysql.Close()

	//Gin
	Routes.InitGin()
}
