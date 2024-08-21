package repository

import (
	"errors"

	"nftPlantform/api"
	"nftPlantform/models"
	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewGormTransactionRepository(db *gorm.DB) api.TransactionRepository {
	return &GormTransactionRepository{db: db}
}

func (r *GormTransactionRepository) CreateTransaction(orderID uint, txHash string, amount, gasFee float64) (uint, error) {
	transaction := models.Transaction{
		OrderID: orderID,
		TxHash:  txHash,
		Amount:  amount,
		GasFee:  gasFee,
		Status:  "PENDING",
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
			return nil, errors.New("Transaction not found")
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
			return nil, errors.New("Transaction not found")
		}
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *GormTransactionRepository) UpdateTransactionStatus(id uint, status string) error {
	return r.db.Model(&models.Transaction{}).Where("id = ?", id).Update("status", status).Error
}

func (r *GormTransactionRepository) GetTransactionsByOrderID(orderID uint) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	result := r.db.Where("order_id = ?", orderID).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}