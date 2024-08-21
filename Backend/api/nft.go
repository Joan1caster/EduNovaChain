package api

import (
	"nftPlantform/models"
)

type NFTRepository interface {
	CreateNFT(tokenID, contractAddress string, ownerID, creatorID int64, metadataURI string) (int64, error)
	GetNFTByID(id int64) (*models.NFT, error)
	GetNFTByTokenID(tokenID string) (*models.NFT, error)
	UpdateNFT(nft *models.NFT) error
	DeleteNFT(id int64) error
	GetNFTsByOwnerID(ownerID int64) ([]*models.NFT, error)
}