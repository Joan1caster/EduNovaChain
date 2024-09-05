package repository

import (
	"errors"
	"sort"

	"gorm.io/gorm"

	_"nftPlantform/api"
	"nftPlantform/models"
	"nftPlantform/utils"
)

type GormNFTRepository struct {
	db *gorm.DB
}

func NewGormNFTRepository(db *gorm.DB) *GormNFTRepository {
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

func (r *GormNFTRepository) GetLatestNFT(number uint) (*[]models.NFT, error) {
	var nfts []models.NFT
	result := r.db.Order("created_at DESC").Limit(3).Find(&nfts)
    if result.Error != nil {
        return nil, result.Error
    }
    return &nfts, nil
}

func (r *GormNFTRepository) GetFavoriteTopic(userID uint) (*models.Topic, error) {
	var topics []models.Topic
	err := r.db.Table("topics").
		Select("topics.id, topics.name, user_topic_visits.visit_count").
		Joins("INNER JOIN user_topic_visits ON topics.id = user_topic_visits.topic_id").
		Where("user_topic_visits.user_id = ?", userID).
		Order("user_topic_visits.visit_count DESC").
		Limit(1).
		Find(&topics).Error

	return &topics[0], err
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

func (r *GormNFTRepository) GetNFTsByOwnerID(ownerID uint) ([]*models.NFT, error) {
	var nfts []*models.NFT
	result := r.db.Where("owner_id = ?", ownerID).Find(&nfts)
	if result.Error != nil {
		return nil, result.Error
	}
	return nfts, nil
}

func (r *GormNFTRepository) GetNFTsByCreatorID(creatorID uint) ([]*models.NFT, error) {
	var nfts []*models.NFT

	result := r.db.Where("creator_id = ?", creatorID).
		Preload("Owner").
		Preload("Creator").
		Order("created_at DESC").
		Limit(100).
		Find(&nfts)

	if result.Error != nil {
		return nil, result.Error
	}

	return nfts, nil
}

func (r *GormNFTRepository) GetMostVisitedNFTsInTopic(topicID uint, limit int) ([]models.NFT, error) {
	var nfts []models.NFT
	err := r.db.Table("nfts").
		Joins("JOIN nft_topics ON nfts.id = nft_topics.nft_id").
		Where("nft_topics.topic_id = ?", topicID).
		Order("nfts.view_count DESC").  // 假设 NFT 模型有 view_count 字段
		Limit(limit).
		Find(&nfts).Error
	return nfts, err
}

func (r *GormNFTRepository) GetNFTByClassification(classification string) ([]*models.NFT, error) {
	var nfts []*models.NFT

	result := r.db.Where("classification = ?", classification).
		Limit(100).
		Find(&nfts)

	if result.Error != nil {
		return nil, result.Error
	}

	return nfts, nil
}

func (r *GormNFTRepository) GetNFTByFeature(feature *[512]float32, similarityThreshold float32) ([]*models.NFTWithSimilarity, error) {
	var result []*models.NFTWithSimilarity
	var offset int
	var lastID uint = 0
	var batchSize int = 1000
	for {
		var batchNFTs []models.NFT
		if err := r.db.Select("id, SummaryFeature").
			Where("id > ?", lastID).
			Order("id").
			Limit(batchSize).
			Find(&batchNFTs).Error; err != nil {
			return nil, err
		}

		if len(batchNFTs) == 0 || len(result) > 100{
			break
		}

		for _, nft := range batchNFTs {
			summaryFeatureFloat, err := utils.BlobToFloat32Array(nft.SummaryFeature)
			if err != nil {
				return nil, err
			}
			similarity := utils.CalculateSimilarity(feature, &summaryFeatureFloat)

			if similarity >= similarityThreshold {
				result = append(result, &models.NFTWithSimilarity{
					NFT:        &nft,
					Similarity: similarity,
				})
			}
		}

		offset += batchSize
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Similarity > result[j].Similarity
	})
	return result, nil
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


func (r *GormNFTRepository) UpdateNFT(nft *models.NFT) error {
	return r.db.Save(nft).Error
}

func (r *GormNFTRepository) DeleteNFT(id uint) error {
	return r.db.Delete(&models.NFT{}, id).Error
}