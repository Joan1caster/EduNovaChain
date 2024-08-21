package repository

import (
	"errors"

	"nftPlantform/api"
	"nftPlantform/models"
	"gorm.io/gorm"
)

type GormNFTRepository struct {
	db *gorm.DB
}

func NewGormNFTRepository(db *gorm.DB) api.NFTRepository {
	return &GormNFTRepository{db: db}
}

func (r *GormNFTRepository) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string) (uint, error) {
	nft := models.NFT{
		TokenID:         tokenID,
		ContractAddress: contractAddress,
		OwnerID:         ownerID,
		CreatorID:       creatorID,
		MetadataURI:     metadataURI,
	}
	result := r.db.Create(&nft)
	if result.Error != nil {
		return 0, result.Error
	}
	return nft.ID, nil
}

func (r *GormNFTRepository) GetNFTByID(id uint) (*models.NFT, error) {
	var nft models.NFT
	result := r.db.First(&nft, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("NFT not found")
		}
		return nil, result.Error
	}
	return &nft, nil
}

func (r *GormNFTRepository) GetNFTByTokenID(tokenID string) (*models.NFT, error) {
	var nft models.NFT
	result := r.db.Where("token_id = ?", tokenID).First(&nft)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("NFT not found")
		}
		return nil, result.Error
	}
	return &nft, nil
}

func (r *GormNFTRepository) UpdateNFT(nft *models.NFT) error {
	return r.db.Save(nft).Error
}

func (r *GormNFTRepository) DeleteNFT(id uint) error {
	return r.db.Delete(&models.NFT{}, id).Error
}

func (r *GormNFTRepository) GetNFTsByOwnerID(ownerID uint) ([]*models.NFT, error) {
	var nfts []*models.NFT
	result := r.db.Where("owner_id = ?", ownerID).Find(&nfts)
	if result.Error != nil {
		return nil, result.Error
	}
	return nfts, nil
}