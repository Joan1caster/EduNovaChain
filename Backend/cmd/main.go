package main

import (
	"log"
	"nftPlantform/config"
	routes "nftPlantform/router"

	"nftPlantform/internal/database"
	_ "nftPlantform/service"
)

func main() {
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
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
