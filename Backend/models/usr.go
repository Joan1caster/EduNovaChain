package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique;not null"`
	Email         string `gorm:"unique;not null"`
	PasswordHash  string `gorm:"not null"`
	WalletAddress string `gorm:"unique;not null"`
	Topics        []Topic `gorm:"many2many:user_topics;"` 
	OwnedNFTs     []NFT   `gorm:"foreignKey:OwnerID"`
	CreatedNFTs   []NFT   `gorm:"foreignKey:CreatorID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Claims struct {
	Wallet string `json:"wallet"`
	UUID   string `json:"uuid"`
	jwt.RegisteredClaims
}
