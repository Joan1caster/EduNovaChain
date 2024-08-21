package models

import (
	"time"
)

type NFT struct {
	ID              int64
	TokenID         string
	ContractAddress string
	OwnerID         int64
	CreatorID       int64
	MetadataURI     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}