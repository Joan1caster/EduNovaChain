package models

import (
	"time"
	"database/sql"
)

type Order struct {
	ID        int64
	SellerID  int64
	BuyerID   sql.NullInt64
	NFTID     int64
	Price     float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

