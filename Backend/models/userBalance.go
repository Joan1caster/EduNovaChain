package models

import (
	"time"
)

type UserBalance struct {
	UserID    int64
	Balance   float64
	UpdatedAt time.Time
}