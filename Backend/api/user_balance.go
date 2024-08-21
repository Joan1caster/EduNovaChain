package api

import (
	_ "nftPlantform/models"
)

type UserBalanceRepository interface {
	GetBalance(userID uint) (float64, error)
	UpdateBalance(userID uint, newBalance float64) error
	AddToBalance(userID uint, amount float64) error
	SubtractFromBalance(userID uint, amount float64) error
}
