package repository

import (
	"errors"
	"gorm.io/gorm"

	"nftPlantform/api"
	"nftPlantform/models"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) api.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) CreateUser(username, email, passwordHash, walletAddress string) (uint, error) {
	user := models.User{
		Username:      username,
		Email:         email,
		PasswordHash:  passwordHash,
		WalletAddress: walletAddress,
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *GormUserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) GetUserByWalletAddress(walletAddress string) (*models.User, error) {
	var user models.User
	result := r.db.Where("wallet_address = ?", walletAddress).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *GormUserRepository) UpdateUser(id uint, updates map[string]interface{}) error {
	result := r.db.Model(&models.User{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
