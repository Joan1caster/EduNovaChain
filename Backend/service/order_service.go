package service

import (
	"errors"

	"nftPlantform/api"
)

type OrderService struct {
	nftRepo   api.NFTRepository
	orderRepo api.OrderRepository
}

func NewOrderService(
	nftRepo api.NFTRepository,
	orderRepo api.OrderRepository,
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

	return s.orderRepo.CreateOrder(sellerID, nftID, price)
}

func (s *OrderService) ValidateOrderStatus(orderID uint) error {
	if s.orderRepo == nil {
		return errors.New("order repository is not initialized")
	}

	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return errors.New("error fetching order")
	}
	if order == nil {
		return errors.New("order not found")
	}
	if order.Status == "OPEN" {
		return nil
	}
	return errors.New("order is not open")
}
