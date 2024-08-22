package api

import (
	"nftPlantform/models"
)

type TransactionRepository interface {
	CreateTransaction(orderID uint, txHash, amount, gasFee, status string) (uint, error)
	GetTransactionByID(id uint) (*models.Transaction, error)
	GetTransactionByTxHash(txHash string) (*models.Transaction, error)
	UpdateTransactionStatus(id uint, status string) error
	GetTransactionsByOrderID(orderID uint) ([]*models.Transaction, error)
	UpdateTransactionGasFee(orderID uint, gasFeeEther string) error
}
