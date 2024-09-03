package service

import (
	"errors"

	"nftPlantform/api"
	"nftPlantform/models"
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

func (s *OrderService) ValidateOrderStatus(orderID uint, sellerWallet string) error {
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
	if order.Seller.WalletAddress != sellerWallet {
		return errors.New("order not build by you")
	}
	if order.Status == "OPEN" {
		return nil
	}
	return errors.New("order is not open")
}

func (s OrderService) GetOrderByID(orderID uint) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s OrderService) CancelOrder(orderID uint) error {
	err := s.orderRepo.CancelOrder(orderID)
	if err != nil {
		return err
	}
	return nil
}
