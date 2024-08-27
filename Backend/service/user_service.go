package service

import (
	"errors"
	"nftPlantform/api"
	"nftPlantform/models"
	"nftPlantform/utils"
)

type UserService struct {
	userRepo api.UserRepository
}

func (s *UserService) AuthenticateUser(address, signature, message string) (*models.User, error) {
	// Verify the signature
	if !utils.VerifySignature(address, signature, message) {
		return nil, errors.New("invalid signature")
	}

	// Check if the user exists, if not, create a new user
	user, err := s.userRepo.GetUserByWalletAddress(address)
	if err != nil {
		// User doesn't exist, create a new one
		userID, err := s.userRepo.CreateUser("", "", "", address)
		if err != nil {
			return nil, err
		}
		user, err = s.userRepo.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
