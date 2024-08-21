package main

import (
	"log"
	_ "nftPlantform/models"
	"nftPlantform/repository"

	"gorm.io/gorm"
)

func TestUser(Db *gorm.DB) {
	userRepo := repository.NewGormUserRepository(Db)

	// initialUser := &models.User{
	// 	Username:      "testuser",
	// 	Email:         "test@example.com",
	// 	PasswordHash:  "oldhash",
	// 	WalletAddress: "0x1234567890123456789012345678901234567890",
	// }

	// err := Db.Create(initialUser).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	user, err := userRepo.GetUserByUsername("玉皇大帝")
	if err != nil {
		log.Fatal(err)
	}

	updates := map[string]interface{}{
		"username": "玉皇大帝",
	}

	err = userRepo.UpdateUser(user.ID, updates)
	if err != nil {
		log.Fatal(err)
	}
	userSelect, err := userRepo.GetUserByID(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userSelect)
}
