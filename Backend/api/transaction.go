package api

import (
	"nftPlantform/models"
)

type TransactionRepository interface {
	CreateTransaction(orderID uint, txHash string, amount, gasFee float64) (uint, error)
	GetTransactionByID(id uint) (*models.Transaction, error)
	GetTransactionByTxHash(txHash string) (*models.Transaction, error)
	UpdateTransactionStatus(id uint, status string) error
	GetTransactionsByOrderID(orderID uint) ([]*models.Transaction, error)
}
