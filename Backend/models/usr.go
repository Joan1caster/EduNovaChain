package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID            uint    `gorm:"primaryKey"`
	Username      string  `gorm:"type:varchar(30);"`
	Email         string  `gorm:"type:varchar(30);"`
	PasswordHash  string  `gorm:"type:varchar(18);"`
	WalletAddress string  `gorm:"unique;not null"`
	Topics        []Topic `gorm:"many2many:user_topics;"`
	OwnedNFTs     []NFT   `gorm:"foreignKey:OwnerID"`
	CreatedNFTs   []NFT   `gorm:"foreignKey:CreatorID"`
	Likes         []Like  `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Claims struct {
	UserID uint   `json:"userid"`
	Wallet string `json:"wallet"`
	UUID   string `json:"uuid"`
	jwt.RegisteredClaims
}
