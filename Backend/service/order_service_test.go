package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"nftPlantform/api"
	"nftPlantform/models"
)

type MockOrderRepository struct {
	mock.Mock
	api.OrderRepository
}

func (m *MockOrderRepository) ListNFTForSale(sellerID, nftID uint, price float64) (uint, error) {
	args := m.Called(sellerID, nftID, price)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockOrderRepository) ValidateOrderStatus(orderID uint) error {
	args := m.Called(orderID)
	return args.Error(0)
}

func (m *MockOrderRepository) CreateOrder(sellerID, nftID uint, price float64) (uint, error) {
	args := m.Called(sellerID, nftID, price)
	return args.Get(0).(uint), args.Error(1)
}

func TestListNFT(t *testing.T) {
	mockOrderRepo := new(MockOrderRepository)
	mockNFTRepo := new(MockNFTRepository)
	OrderService := NewOrderService(mockNFTRepo, mockOrderRepo)
	mockNFTRepo.On("GetNFTByID", uint(1)).Return(&models.NFT{ID: 1, OwnerID: 1}, nil)
	mockOrderRepo.On("CreateOrder", uint(1), uint(1), float64(1)).Return(uint(1), nil)
	mockOrderRepo.On("ListNFTForSale", uint(1), uint(1), float64(1))
	orderID, err := OrderService.ListNFTForSale(uint(1), uint(1), float64(1))
	assert.NoError(t, err)
	assert.Equal(t, orderID, uint(1))
}

func TestValidateOrderStatus(t *testing.T) {
    mockOrderRepo := new(MockOrderRepository)
    orderService := NewOrderService(nil, mockOrderRepo)  // 确保正确初始化

    orderID := uint(100)

    // 模拟 GetOrderByID 的行为，返回一个有效的订单
    mockOrderRepo.On("GetOrderByID", orderID).Return(&models.Order{ID: orderID, Status: "OPEN"}, nil)

    err := orderService.ValidateOrderStatus(orderID)
    assert.NoError(t, err)

    mockOrderRepo.AssertExpectations(t)
}