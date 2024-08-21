package api

import (
	_"nftPlantform/models"
)

type UserBalanceRepository interface {
	GetBalance(userID int64) (float64, error)
	UpdateBalance(userID int64, newBalance float64) error
	AddToBalance(userID int64, amount float64) error
	SubtractFromBalance(userID int64, amount float64) error
}