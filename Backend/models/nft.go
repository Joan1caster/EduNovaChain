package models

import (
	"time"
)

type NFT struct {
	ID              uint   `gorm:"primaryKey"`
	TokenID         string `gorm:"unique;not null"`
	ContractAddress string `gorm:"not null"`
	OwnerID         uint   `gorm:"not null"`
	CreatorID       uint   `gorm:"not null"`
	Classification  string
	MetadataURI     string `gorm:"not null"`
	SummaryFeature  []byte `gorm:"type:blob;not null"`
	ContentFeature  []byte `gorm:"type:blob;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Owner           User `gorm:"foreignKey:OwnerID"`
	Creator         User `gorm:"foreignKey:CreatorID"`
}

type NFTWithSimilarity struct {
	*NFT
	Similarity float32
}
