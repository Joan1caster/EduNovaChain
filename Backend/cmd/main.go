package main

import (
	"log"
	"nftPlantform/api"
	"nftPlantform/internal/database"
	"nftPlantform/repository"
	_ "nftPlantform/service"
)

func main() {
	database.ConnectDB()
	// defer database.CloseDB()
	db := database.GetDB()
	if db == nil {
		log.Println("database open failed")
	}
	var userRepo api.UserRepository = repository.NewMySQLUserRepository(db)

	// 使用userRepo的方法
	// newUserID, err := userRepo.CreateUser("bill_doe", "john@example1.com", "hashed_password1", "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92267")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(newUserID)
	user, err := userRepo.GetUserByID(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}
