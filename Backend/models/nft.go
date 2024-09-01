package models

import (
	"time"
)

type NFT struct {
	ID              uint         `gorm:"primaryKey"`
	TokenID         string       `gorm:"unique;not null"`
	ContractAddress string       `gorm:"not null"`
	OwnerID         uint         `gorm:"not null"`
	CreatorID       uint         `gorm:"not null"`
	MetadataURI     string       `gorm:"not null"`
	AbstractFeature [512]float32 `gorm:"type:float[];not null"`
	MetadataFeature [512]float32 `gorm:"type:float[];not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Owner           User `gorm:"foreignKey:OwnerID"`
	Creator         User `gorm:"foreignKey:CreatorID"`
}
