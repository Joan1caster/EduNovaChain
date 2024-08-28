package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique;not null"`
	Email         string `gorm:"unique;not null"`
	PasswordHash  string `gorm:"not null"`
	WalletAddress string `gorm:"unique;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Claims struct {
	Wallet string `json:"wallet"`
	UUID   string `json:"uuid"`
	jwt.RegisteredClaims
}
