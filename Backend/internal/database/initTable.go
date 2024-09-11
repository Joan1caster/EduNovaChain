package database

import (
	"log"
	"nftPlantform/models"
)

func InitTable() {
	db := GetDB()
	err := db.AutoMigrate(&models.User{}, &models.NFT{}, &models.Order{}, &models.Transaction{}, &models.Like{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database tables created successfully")
}
