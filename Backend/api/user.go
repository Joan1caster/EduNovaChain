package api

import (
	"nftPlantform/models"
)

type UserRepository interface {
	CreateUser(username, email, passwordHash, walletAddress string) (int64, error)
	GetUserByID(id int64) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByWalletAddress(walletAddress string) (*models.User, error)
	UpdateUser(id uint, updates map[string]interface{}) error
	DeleteUser(id int64) error
}
