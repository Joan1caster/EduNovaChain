package main

import (
	"log"

	"gorm.io/gorm"

	"nftPlantform/config"
	"nftPlantform/internal/database"
	_ "nftPlantform/service"
)

var Db *gorm.DB

func main() {
	cfg := config.LoadConfig()
	log.Println("connect rpc-url:", cfg.Contract.Eth_rpc_url)
	database.ConnectDB()
	defer database.CloseDB()
	database.InitTable()
	Db = database.GetDB()
	TestUser(Db)
}
