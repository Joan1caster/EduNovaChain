package models

import (
	"time"
)

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	SellerID  uint      `gorm:"not null"`
	BuyerID   *uint
	NFTID     uint      `gorm:"not null"`
	Price     float64   `gorm:"type:decimal(20,8);not null"`
	Status    string    `gorm:"type:enum('OPEN','COMPLETED','CANCELLED');not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Seller    User      `gorm:"foreignKey:SellerID"`
	Buyer     *User     `gorm:"foreignKey:BuyerID"`
	NFT       NFT       `gorm:"foreignKey:NFTID"`
	
}

