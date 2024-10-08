package service

import (
	"errors"

	"nftPlantform/models"
	"nftPlantform/repository"
)

type OrderService struct {
	nftRepo   *repository.GormNFTRepository
	orderRepo *repository.GormOrderRepository
}

func NewOrderService(
	nftRepo *repository.GormNFTRepository,
	orderRepo *repository.GormOrderRepository,
) *OrderService {

	return &OrderService{
		nftRepo:   nftRepo,
		orderRepo: orderRepo,
	}
}

func (s *OrderService) ListNFTForSale(sellerID, nftID uint, price float64) (uint, error) {
	nft, err := s.nftRepo.GetNFTByID(nftID)

	if err != nil {
		return 0, err
	}

	if nft.OwnerID != sellerID {
		return 0, errors.New("you are not owner")
	}

	order, err := s.orderRepo.GetOpenOrdersBySellerIDAndNFTID(sellerID, nftID)
	if err != nil {
		return 0, err
	}
	if order != nil && order.Status == "OPEN" {
		if order.Price != price {
			order.Price = price
			err := s.orderRepo.UpdateOrder(order)
			if err != nil {
				return 0, err
			}
			return order.ID, nil
		} else {
			return 0, errors.New("order has been submit, please do not repeat")
		}
	}

	return s.orderRepo.CreateOrder(sellerID, nftID, price)
}

func (s *OrderService) ValidateOrderStatus(orderID uint, userID uint) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return errors.New("error fetching order")
	}
	if order == nil {
		return errors.New("order not found")
	}
	if order.SellerID != userID {
		return errors.New("order not build by you")
	}
	if order.Status == "OPEN" {
		return nil
	}
	return errors.New("order is not open")
}

func (s *OrderService) ValidateOrderIsconfirmed(orderID uint) (bool, error) {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return false, errors.New("error fetching order")
	}
	if order == nil {
		return false, errors.New("order not found")
	}
	if order.Status == "COMPLETED" {
		return true, nil
	}
	if order.Status == "CANCELLED" {
		return false, errors.New("order has been canceled")
	}
	return false, nil
}

func (s OrderService) GetOrderByID(orderID uint) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s OrderService) GetCompletedOrdersByNFTID(nftId uint) ([]*models.Order, error) {
	var orders []*models.Order
	orders, err := s.orderRepo.GetCompletedOrdersByNFTID(nftId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s OrderService) CancelOrder(orderID uint) error {
	err := s.orderRepo.CancelOrder(orderID)
	if err != nil {
		return err
	}
	return nil
}
