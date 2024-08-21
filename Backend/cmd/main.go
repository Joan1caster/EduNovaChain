package main

import (
	"nftPlantform/internal/database"
	_ "nftPlantform/service"

	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {
	database.ConnectDB()
	defer database.CloseDB()
	database.InitTable()
	Db = database.GetDB()
	TestUser(Db)
}
