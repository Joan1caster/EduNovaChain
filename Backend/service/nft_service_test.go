package service

import (
	"nftPlantform/api"
	"nftPlantform/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNFTRepository is a mock of NFTRepository
type MockNFTRepository struct {
	mock.Mock
	api.NFTRepository
}

func (m *MockNFTRepository) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, abstractFeature, metadataFeature [512]float32) (uint, error) {
	args := m.Called(tokenID, contractAddress, ownerID, creatorID, metadataURI, abstractFeature, metadataFeature )
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockNFTRepository) GetNFTByID(id uint) (*models.NFT, error) {
	args := m.Called(id)
	return args.Get(0).(*models.NFT), args.Error(1)
}

func (m *MockNFTRepository) GetNFTsByOwnerID(ownerID uint) ([]*models.NFT, error) {
	args := m.Called(ownerID)
	return args.Get(0).([]*models.NFT), args.Error(1)
}

func (m *MockNFTRepository) UpdateNFT(nft *models.NFT) error {
	args := m.Called(nft)
	return args.Error(0)
}

func TestCreateNFT(t *testing.T) {
	mockNFTRepo := new(MockNFTRepository)
	service := NewNFTService(mockNFTRepo)

	mockNFTRepo.On("CreateNFT", "token123", "0x123", uint(1), uint(2), "uri://metadata").Return(uint(1), nil)
	var abstractFeature [512]float32
	var metadataFeature [512]float32
	id, err := service.CreateNFT("token123", "0x123", 1, 2, "uri://metadata", abstractFeature, metadataFeature)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), id)
	mockNFTRepo.AssertExpectations(t)
}

func TestGetNFTDetails(t *testing.T) {
	mockNFTRepo := new(MockNFTRepository)
	service := NewNFTService(mockNFTRepo)

	expectedNFT := &models.NFT{ID: 1, TokenID: "token123"}
	mockNFTRepo.On("GetNFTByID", uint(1)).Return(expectedNFT, nil)

	nft, err := service.GetNFTDetails(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedNFT, nft)
	mockNFTRepo.AssertExpectations(t)
}

func TestTransferNFT(t *testing.T) {
	mockNFTRepo := new(MockNFTRepository)
	service := NewNFTService(mockNFTRepo)

	nft := &models.NFT{ID: 1, OwnerID: 1}
	mockNFTRepo.On("GetNFTByID", uint(1)).Return(nft, nil)
	mockNFTRepo.On("UpdateNFT", mock.AnythingOfType("*models.NFT")).Return(nil)

	err := service.TransferNFT(1, 1, 2)

	assert.NoError(t, err)
	assert.Equal(t, uint(2), nft.OwnerID)
	mockNFTRepo.AssertExpectations(t)
}

func TestTransferNFT_NotOwner(t *testing.T) {
	mockNFTRepo := new(MockNFTRepository)
	service := NewNFTService(mockNFTRepo)

	nft := &models.NFT{ID: 1, OwnerID: 2}
	mockNFTRepo.On("GetNFTByID", uint(1)).Return(nft, nil)

	err := service.TransferNFT(1, 1, 3)

	assert.Error(t, err)
	assert.Equal(t, "user does not own this NFT", err.Error())
	mockNFTRepo.AssertExpectations(t)
}
