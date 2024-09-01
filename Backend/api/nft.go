package api

import (
	"nftPlantform/models"
)

type NFTRepository interface {
	CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, abstractFeature, metadataFeature [512]float32) (uint, error)
	GetNFTByID(id uint) (*models.NFT, error)
	GetNFTByTokenID(tokenID string) (*models.NFT, error)
	UpdateNFT(nft *models.NFT) error
	DeleteNFT(id uint) error
	GetNFTsByOwnerID(ownerID uint) ([]*models.NFT, error)
	GetAbstractFeatures(batchSize int) ([][512]float32, error)
	GetMetadataFeatures(batchSize int) ([][512]float32, error)
}
