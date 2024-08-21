package repository

import (
	"errors"

	"nftPlantform/api"
	"nftPlantform/models"
	"gorm.io/gorm"
)

type GormUserBalanceRepository struct {
	db *gorm.DB
}

func NewGormUserBalanceRepository(db *gorm.DB) api.UserBalanceRepository {
	return &GormUserBalanceRepository{db: db}
}

func (r *GormUserBalanceRepository) GetBalance(userID uint) (float64, error) {
	var userBalance models.UserBalance
	result := r.db.Where("user_id = ?", userID).First(&userBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil // 如果记录不存在，返回余额为0
		}
		return 0, result.Error
	}
	return userBalance.Balance, nil
}

func (r *GormUserBalanceRepository) UpdateBalance(userID uint, newBalance float64) error {
	return r.db.Model(&models.UserBalance{}).Where("user_id = ?", userID).UpdateColumn("balance", newBalance).Error
}

func (r *GormUserBalanceRepository) AddToBalance(userID uint, amount float64) error {
	return r.db.Model(&models.UserBalance{}).Where("user_id = ?", userID).UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error
}

func (r *GormUserBalanceRepository) SubtractFromBalance(userID uint, amount float64) error {
	return r.db.Model(&models.UserBalance{}).Where("user_id = ?", userID).UpdateColumn("balance", gorm.Expr("balance - ?", amount)).Error
}
