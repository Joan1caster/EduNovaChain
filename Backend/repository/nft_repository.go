package repository

import (
	"errors"

	"gorm.io/gorm"

	"nftPlantform/api"
	"nftPlantform/models"
)

type GormNFTRepository struct {
	db *gorm.DB
}

func NewGormNFTRepository(db *gorm.DB) api.NFTRepository {
	return &GormNFTRepository{db: db}
}

func (r *GormNFTRepository) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, abstractFeature, metadataFeature [512]float32) (uint, error) {
	nft := models.NFT{
		TokenID:         tokenID,
		ContractAddress: contractAddress,
		OwnerID:         ownerID,
		CreatorID:       creatorID,
		MetadataURI:     metadataURI,
		AbstractFeature: abstractFeature,
		MetadataFeature: metadataFeature,
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

func (r *GormNFTRepository) GetAbstractFeatures(batchSize int) ([][512]float32, error) {
	var allAbstractFeatures [][512]float32
	var lastID uint = 0
	for {
		var batch []models.NFT
		if err := r.db.Select("id, AbstractFeature").
			Where("id > ?", lastID).
			Order("id").
			Limit(batchSize).
			Find(&batch).Error; err != nil {
			return nil, err
		}

		if len(batch) == 0 {
			break
		}

		for _, nft := range batch {
			allAbstractFeatures = append(allAbstractFeatures, nft.AbstractFeature)
			lastID = nft.ID
		}

		if len(batch) < batchSize {
			break
		}
	}
	return allAbstractFeatures, nil
}

func (r *GormNFTRepository) GetMetadataFeatures(batchSize int) ([][512]float32, error) {
	var allMetadatatFeatures [][512]float32
	var lastID uint = 0
	for {
		var batch []models.NFT
		if err := r.db.Select("id, MetadataFeature").
			Where("id > ?", lastID).
			Order("id").
			Limit(batchSize).
			Find(&batch).Error; err != nil {
			return nil, err
		}

		if len(batch) == 0 {
			break
		}

		for _, nft := range batch {
			allMetadatatFeatures = append(allMetadatatFeatures, nft.AbstractFeature)
			lastID = nft.ID
		}

		if len(batch) < batchSize {
			break
		}
	}
	return allMetadatatFeatures, nil
}
