package repository

import (
	"errors"

	"nftPlantform/models"

	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewGormTransactionRepository(db *gorm.DB) *GormTransactionRepository {
	return &GormTransactionRepository{db: db}
}

func (r *GormTransactionRepository) CreateTransaction(orderID uint, txHash, amount, gasFee, status string) (uint, error) {
	transaction := models.Transaction{
		OrderID:     orderID,
		TxHash:      txHash,
		Amount:      amount,
		GasFeeEther: gasFee,
		Status:      status,
	}
	result := r.db.Create(&transaction)
	if result.Error != nil {
		return 0, result.Error
	}
	return transaction.ID, nil
}

func (r *GormTransactionRepository) GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	result := r.db.First(&transaction, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *GormTransactionRepository) GetTransactionByTxHash(txHash string) (*models.Transaction, error) {
	var transaction models.Transaction
	result := r.db.Where("tx_hash = ?", txHash).First(&transaction)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *GormTransactionRepository) UpdateTransactionStatus(orderID uint, status string) error {
	return r.db.Model(&models.Transaction{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *GormTransactionRepository) UpdateTransactionGasFee(orderID uint, gasFeeEther string) error {
	return r.db.Model(&models.Transaction{}).Where("id = ?", orderID).Update("GasFee", gasFeeEther).Error
}

func (r *GormTransactionRepository) GetTransactionsByOrderID(orderID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	result := r.db.Where("order_id = ?", orderID).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
