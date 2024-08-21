package api

import (
	"nftPlantform/models"
)

type UserRepository interface {
	CreateUser(username, email, passwordHash, walletAddress string) (uint, error)
	GetUserByID(id uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByWalletAddress(walletAddress string) (*models.User, error)
	UpdateUser(id uint, updates map[string]interface{}) error
	DeleteUser(id uint) error
}
