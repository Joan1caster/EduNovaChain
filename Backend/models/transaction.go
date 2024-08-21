package models

import (
	"time"
)

type Transaction struct {
	ID        int64
	OrderID   int64
	TxHash    string
	Amount    float64
	GasFee    float64
	Status    string
	CreatedAt time.Time
}
