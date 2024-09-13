package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"

	"nftPlantform/config"
	"nftPlantform/handlers"
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
	router.Use(handlers.ErrorHandler())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},        // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	err = router.Run("127.0.0.1:4457")
	if err != nil {
		return
	}
}
