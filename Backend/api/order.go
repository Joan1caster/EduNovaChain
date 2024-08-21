package api

import (
	"nftPlantform/models"
)

type OrderRepository interface {
	CreateOrder(sellerID, nftID int64, price float64) (int64, error)
	GetOrderByID(id int64) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id int64) error
	GetOpenOrdersByNFTID(nftID int64) ([]*models.Order, error)
	CompleteOrder(id int64, buyerID int64) error
	CancelOrder(id int64) error
}