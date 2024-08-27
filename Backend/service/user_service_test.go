package service

import (
	_ "testing"

	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"nftPlantform/api"
	"nftPlantform/models"
)

type MockUserRepository struct {
	mock.Mock
	api.UserRepository
}

func (m *MockUserRepository) CreateUser(username, email, passwordHash, walletAddress string) (uint, error) {
	args := m.Called(username, email, passwordHash, walletAddress)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(id uint) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByWalletAddress(walletAddress string) (*models.User, error) {
	args := m.Called(walletAddress)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(id uint, updates map[string]interface{}) error {
	args := m.Called(id, updates)
	return args.Error(0)
}
