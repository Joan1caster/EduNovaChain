package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"nftPlantform/config"
	routes "nftPlantform/router"

	"nftPlantform/internal/database"
	"nftPlantform/utils"
)

func main() {
	utils.SetupLogger()
	err := config.LoadConfig("../config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration")
	}
	log.Println("connect rpc-url:", config.AppConfig.Contract.Eth_rpc_url)
	database.ConnectDB()
	defer database.CloseDB()
	database.InitTable()
	db := database.GetDB()
	router := routes.SetupRouter(db)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))
	err = router.Run("127.0.0.1:4455")
	if err != nil {
		return
	}
}
