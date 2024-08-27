package repository

import (
	"errors"

	"nftPlantform/api"
	"nftPlantform/models"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) api.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) CreateOrder(sellerID, nftID uint, price float64) (uint, error) {
	order := models.Order{
		SellerID: sellerID,
		NFTID:    nftID,
		Price:    price,
		Status:   "OPEN",
	}
	result := r.db.Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return order.ID, nil
}

func (r *GormOrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	if r.db == nil {
		return nil, errors.New("Database connection not initialized")
	}

	var order models.Order
	result := r.db.Take(&order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		// 可以在这里添加日志
		return nil, errors.New("failed to get order")
	}
	return &order, nil
}

func (r *GormOrderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *GormOrderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}

func (r *GormOrderRepository) GetOpenOrdersByNFTID(nftID uint) ([]*models.Order, error) {
	var orders []*models.Order
	result := r.db.Where("nft_id = ? AND status = ?", nftID, "OPEN").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (r *GormOrderRepository) CompleteOrder(id uint, buyerID uint) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Updates(map[string]interface{}{
		"buyer_id": buyerID,
		"status":   "COMPLETED",
	}).Error
}

func (r *GormOrderRepository) CancelOrder(id uint) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", "CANCELLED").Error
}

func (r *GormOrderRepository) ReopenOrder(id uint) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", "OPEN").Error
}
