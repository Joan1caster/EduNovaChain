package api

import (
	"nftPlantform/models"
)

type OrderRepository interface {
	CreateOrder(sellerID, nftID uint, price float64) (uint, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
	GetOpenOrdersByNFTID(nftID uint) ([]*models.Order, error)
	GetCompletedOrdersByNFTID(nftID uint) ([]*models.Order, error)
	CompleteOrder(id uint, buyerID uint) error
	CancelOrder(id uint) error
	ReopenOrder(id uint) error
}
