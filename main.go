package main

import (
	"go-api/databases"
	"go-api/routes"
)

func main() {
	//GORM
	databases.NewMysql()

	//Gin Gonic
	routes.NewGin()
}
