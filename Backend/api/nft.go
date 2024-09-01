package api

import (
	"nftPlantform/models"
)

type NFTRepository interface {
	CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, summaryFeature, contantFeature [512]float32) (uint, error)
	GetNFTByID(id uint) (*models.NFT, error)
	GetNFTByTokenID(tokenID string) (*models.NFT, error)
	UpdateNFT(nft *models.NFT) error
	DeleteNFT(id uint) error
	GetNFTsByOwnerID(ownerID uint) ([]*models.NFT, error)
	GetSummaryFeatures(batchSize int) ([][512]float32, error)
	GetContentFeatures(batchSize int) ([][512]float32, error)
}
