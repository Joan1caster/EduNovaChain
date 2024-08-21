package api

import (
	"nftPlantform/models"
)

type TransactionRepository interface {
	CreateTransaction(orderID int64, txHash string, amount, gasFee float64) (int64, error)
	GetTransactionByID(id int64) (*models.Transaction, error)
	GetTransactionByTxHash(txHash string) (*models.Transaction, error)
	UpdateTransactionStatus(id int64, status string) error
	GetTransactionsByOrderID(orderID int64) ([]*models.Transaction, error)
}