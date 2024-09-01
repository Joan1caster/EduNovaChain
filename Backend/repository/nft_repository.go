package repository

import (
	"errors"

	"gorm.io/gorm"

	"nftPlantform/api"
	"nftPlantform/models"
	"nftPlantform/utils"
)

type GormNFTRepository struct {
	db *gorm.DB
}

func NewGormNFTRepository(db *gorm.DB) api.NFTRepository {
	return &GormNFTRepository{db: db}
}

func (r *GormNFTRepository) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, summaryFeature, contentFeature [512]float32) (uint, error) {
	summaryFeatureBlob, err := utils.Float32ArrayToBlob(summaryFeature)
	if err != nil {
		return 0, err
	}
	contentFeatureBlob, err := utils.Float32ArrayToBlob(contentFeature)
	if err != nil {
		return 0, err
	}
	nft := models.NFT{
		TokenID:         tokenID,
		ContractAddress: contractAddress,
		OwnerID:         ownerID,
		CreatorID:       creatorID,
		MetadataURI:     metadataURI,
		SummaryFeature:  summaryFeatureBlob,
		ContentFeature:  contentFeatureBlob,
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

func (r *GormNFTRepository) GetSummaryFeatures(batchSize int) ([][512]float32, error) {
	var allSummaryFeatures [][512]float32
	var lastID uint = 0
	for {
		var batch []models.NFT
		if err := r.db.Select("id, SummaryFeature").
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
			summaryFeatureFloat, err := utils.BlobToFloat32Array(nft.SummaryFeature)
			if err != nil {
				return nil, err
			}
			allSummaryFeatures = append(allSummaryFeatures, summaryFeatureFloat)
			lastID = nft.ID
		}

		if len(batch) < batchSize {
			break
		}
	}
	return allSummaryFeatures, nil
}

func (r *GormNFTRepository) GetContentFeatures(batchSize int) ([][512]float32, error) {
	var allMetadatatFeatures [][512]float32
	var lastID uint = 0
	for {
		var batch []models.NFT
		if err := r.db.Select("id, ContantFeature").
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
			summaryFeatureFloat, err := utils.BlobToFloat32Array(nft.SummaryFeature)
			if err != nil {
				return nil, err
			}
			allMetadatatFeatures = append(allMetadatatFeatures, summaryFeatureFloat)
			lastID = nft.ID
		}

		if len(batch) < batchSize {
			break
		}
	}
	return allMetadatatFeatures, nil
}
