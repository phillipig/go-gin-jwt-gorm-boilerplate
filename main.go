package main

import (
	"go-api/databases"
	"go-api/routes"
)

func main() {
	//GORM
	mysql := databases.NewMysql()
	defer mysql.Close()

	//Gin Gonic
	routes.NewGin()
}
