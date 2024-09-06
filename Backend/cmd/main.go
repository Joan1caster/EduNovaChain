package main

import (
	"log"
	"nftPlantform/config"
	routes "nftPlantform/router"

	"nftPlantform/internal/database"
	"nftPlantform/utils"
)


func main() {
	utils.SetupLogger()
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration")
	}
	log.Println("connect rpc-url:", config.AppConfig.Contract.Eth_rpc_url)
	database.ConnectDB()
	defer database.CloseDB()
	database.InitTable()
	db := database.GetDB()
	router := routes.SetupRouter(db)
	err = router.Run("0.0.0.0:4455")
	if err != nil {
		return
	}
}
