package repository

import (
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"

	_ "nftPlantform/api"
	"nftPlantform/models"
	"nftPlantform/models/dto"
	"nftPlantform/utils"
)

type GormNFTRepository struct {
	db *gorm.DB
}

func NewGormNFTRepository(db *gorm.DB) *GormNFTRepository {
	return &GormNFTRepository{db: db}
}

func (r *GormNFTRepository) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, summaryFeature, contentFeature [512]float32, grade *models.Grade, subject *models.Subject, topic *models.Topic) (uint, error) {
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
		Grades:          []models.Grade{*grade},
		Subjects:        []models.Subject{*subject},
		Topics:          []models.Topic{*topic},
		Categories:      []models.NFTCategory{models.CategoryNewest},
	}
	result := r.db.Create(&nft)
	if result.Error != nil {
		return 0, result.Error
	}
	r.db.Save(nft)
	return nft.ID, nil
}

func (r *GormNFTRepository) GetNFTByID(id uint) (*models.NFT, error) {
	var nft models.NFT
	result := r.db.
		Preload("Grades").
		Preload("Subjects").
		Preload("Topics").
		Preload("Owner").
		Preload("Creator").
		Where("is_for_sale = ?", true).
		First(&nft, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("NFT not found")
		}
		return nil, result.Error
	}

	return &nft, nil
}

func (r *GormNFTRepository) SetNFTUnSaleByID(id uint) error {
	var nft models.NFT
	result := r.db.
		Preload("Grades").
		Preload("Subjects").
		Preload("Topics").
		Preload("Owner").
		Preload("Creator").
		First(&nft, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("NFT not found")
		}
		return result.Error
	}
	nft.IsForSale = false
	if err := r.db.Save(&nft).Error; err != nil { // 使用 Save 方法更新
		return err
	}
	return nil
}

func (r *GormNFTRepository) GetLatestNFT(number uint) (*[]models.NFT, error) {
	var nfts []models.NFT

	result := r.db.Order("created_at DESC").
		Preload("Grades").
		Preload("Subjects").
		Preload("Topics").
		Preload("Owner").
		Preload("Creator").
		Limit(int(number)).
		Find(&nfts)

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
	result := r.db.Where("owner_id = ? AND is_for_sale = ?", ownerID, true).Find(&nfts)
	if result.Error != nil {
		return nil, result.Error
	}
	return nfts, nil
}

func (r *GormNFTRepository) GetNFTsByCreatorID(creatorID uint) ([]*models.NFT, error) {
	var nfts []*models.NFT

	result := r.db.Where("creator_id = ? AND is_for_sale = ?", creatorID, true).
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

func (r *GormNFTRepository) GetNFTByTopicAndType(topicId, typeId *uint, limit uint) (*[]models.NFT, error) {
	var nfts *[]models.NFT

	query := r.db.Model(&models.NFT{}).
		Preload("Grades").
		Preload("Subjects").
		Preload("Topics").
		Preload("Owner").
		Preload("Creator").
		Where("is_for_sale = ?", true) // 添加过滤条件

		// 根据 topicId 筛选
	if topicId != nil {
		query = query.Joins("JOIN nft_topics ON nfts.id = nft_topics.nft_id").
			Where("nft_topics.topic_id = ?", *topicId)
	}

	// 根据 typeId 筛选类别
	if typeId != nil {
		var category models.NFTCategory
		switch *typeId {
		case 1:
			category = models.CategoryNewest
		case 2:
			category = models.CategoryHot
		case 3:
			category = models.CategoryBestseller
		default:
			return nil, fmt.Errorf("无效的 typeId")
		}
		query = query.Where("JSON_CONTAINS(categories, ?)", fmt.Sprintf("\"%s\"", category))
	}

	// 执行查询并限制结果数量
	err := query.Limit(int(limit)).Find(&nfts).Error
	if err != nil {
		return nil, err
	}

	return nfts, nil
}

func (r *GormNFTRepository) GetNFTByDetails(query dto.NFTQuery) ([]*models.NFT, error) {
	var nfts []*models.NFT

	// Start building the query
	tx := r.db.Model(&models.NFT{}).Where("is_for_sale = ?", true) // 添加过滤条件

	// Grade filter
	if query.GradeIDs != nil {
		tx = tx.Joins("JOIN nft_grades ON nfts.id = nft_grades.nft_id").
			Where("nft_grades.grade_id IN ?", *query.GradeIDs)
	}

	// Subjects filter
	if query.Subjects != nil && len(*query.Subjects) > 0 {
		tx = tx.Joins("JOIN nft_subjects ON nfts.id = nft_subjects.nft_id").
			Where("nft_subjects.subject_id IN ?", *query.Subjects)
	}

	// Topics filter
	if query.TopicIds != nil && len(*query.TopicIds) > 0 {
		tx = tx.Joins("JOIN nft_topics ON nfts.id = nft_topics.nft_id").
			Where("nft_topics.topic_id IN ?", *query.TopicIds)
	}
	// Execute the query
	err := tx.Find(&nfts).Error
	if err != nil {
		return nil, err
	}
	return nfts, nil
}

func (r *GormNFTRepository) GetGrade() (*[]dto.IDName, error) {
	var grades *[]dto.IDName

	result := r.db.Model(&models.Grade{}).Select("id", "name").Find(&grades)
	if result.Error != nil {
		return nil, result.Error
	}
	return grades, nil
}

func (r *GormNFTRepository) FindOrCreateSubject(name string) (*models.Subject, error) {
	var subject *models.Subject
	// 尝试查找现有的Subject
	result := r.db.Where("name = ?", name).First(&subject)

	// 如果没有找到,则创建新的Subject
	if result.Error == gorm.ErrRecordNotFound {
		subject = &models.Subject{Name: name}
		if err := r.db.Create(&subject).Error; err != nil {
			return nil, fmt.Errorf("创建Subject失败: %w", err)
		}
	} else if result.Error != nil {
		return nil, fmt.Errorf("查询Subject失败: %w", result.Error)
	}

	return subject, nil
}

func (r *GormNFTRepository) FindOrCreateGrade(name string) (*models.Grade, error) {
	var grade *models.Grade
	// 尝试查找现有的Subject
	result := r.db.Where("name = ?", name).First(&grade)

	// 如果没有找到,则创建新的Subject
	if result.Error == gorm.ErrRecordNotFound {
		grade = &models.Grade{Name: name}
		if err := r.db.Create(&grade).Error; err != nil {
			return nil, fmt.Errorf("创建grade失败: %w", err)
		}
	} else if result.Error != nil {
		return nil, fmt.Errorf("查询grade失败: %w", result.Error)
	}

	return grade, nil
}

func (r *GormNFTRepository) FindOrCreateTopic(name string) (*models.Topic, error) {
	var topic *models.Topic
	// 尝试查找现有的Subject
	result := r.db.Where("name = ?", name).First(&topic)

	// 如果没有找到,则创建新的Subject
	if result.Error == gorm.ErrRecordNotFound {
		topic = &models.Topic{Name: name}
		if err := r.db.Create(&topic).Error; err != nil {
			return nil, fmt.Errorf("创建Subject失败: %w", err)
		}
	} else if result.Error != nil {
		return nil, fmt.Errorf("查询Subject失败: %w", result.Error)
	}

	return topic, nil
}

func (r *GormNFTRepository) GetMostVisitedNFTsInTopic(topicID uint, limit int) ([]models.NFT, error) {
	var nfts []models.NFT
	err := r.db.Table("nfts").
		Joins("JOIN nft_topics ON nfts.id = nft_topics.nft_id").
		Where("nft_topics.topic_id = ?", topicID).
		Order("nfts.view_count DESC"). // 假设 NFT 模型有 view_count 字段
		Limit(limit).
		Find(&nfts).Error
	return nfts, err
}

func (r *GormNFTRepository) GetTopicBySubjectAndGrade(subjectIDs, gradeIDs *[]uint) ([]dto.IDName, error) {
	var results []dto.IDName

	query := r.db.Table("topics").
		Select("DISTINCT topics.id, topics.name").
		Joins("JOIN nft_topics ON topics.id = nft_topics.topic_id").
		Joins("JOIN nfts ON nfts.id = nft_topics.nft_id")

	if subjectIDs != nil && len(*subjectIDs) > 0 {
		query = query.Joins("JOIN nft_subjects ON nfts.id = nft_subjects.nft_id").
			Where("nft_subjects.subject_id IN ?", *subjectIDs)
	}

	if gradeIDs != nil && len(*gradeIDs) > 0 {
		query = query.Joins("JOIN nft_grades ON nfts.id = nft_grades.nft_id").
			Where("nft_grades.grade_id IN ?", *gradeIDs)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GormNFTRepository) GetNFTByClassification(classification string) ([]*models.NFT, error) {
	var nfts []*models.NFT

	result := r.db.Where("classification = ? AND is_for_sale = ?", classification, true).
		Limit(100).
		Find(&nfts)

	if result.Error != nil {
		return nil, result.Error
	}

	return nfts, nil
}

func (r *GormNFTRepository) GetSubjectByGrade(gradeId uint) (*[]models.Subject, error) {
	var subjects *[]models.Subject
	err := r.db.Model(&models.Subject{}).
		Distinct().
		Joins("JOIN nft_subjects ON subjects.id = nft_subjects.subject_id").
		Joins("JOIN nfts ON nfts.id = nft_subjects.nft_id").
		Joins("JOIN nft_grades ON nfts.id = nft_grades.nft_id").
		Where("nft_grades.grade_id = ?", gradeId).
		Find(&subjects).Error
	if err != nil {
		return nil, err
	}
	return subjects, err
}

func (r *GormNFTRepository) GetNFTByFeature(feature *[]float32, similarityThreshold float32) ([]*models.NFTWithSimilarity, error) {
	var result []*models.NFTWithSimilarity
	var offset int
	var lastID uint = 0
	var batchSize int = 1000
	for {
		var batchNFTs []models.NFT
		if err := r.db.Select("id, SummaryFeature").
			Where("id > ?", lastID).
			Where("is_for_sale = ?", true). // 添加过滤条件
			Order("id").
			Limit(batchSize).
			Find(&batchNFTs).Error; err != nil {
			return nil, err
		}

		if len(batchNFTs) == 0 || len(result) > 100 {
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

func (r *GormNFTRepository) GetSummaryFeatures(batchSize int) ([][]float32, error) {
	var allSummaryFeatures [][]float32
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

func (r *GormNFTRepository) LikeNFTByID(userID, nftID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		like := models.Like{UserID: userID, NFTID: nftID}
		err := tx.Create(&like).Error
		if err != nil {
			// 如果是唯一约束错误，我们可以忽略它
			if !errors.Is(err, gorm.ErrDuplicatedKey) {
				return err
			}
			// 如果是重复点赞，我们直接返回，不增加点赞数
			return nil
		}

		// 只有在成功创建Like记录后，才增加NFT的点赞数
		return tx.Model(&models.NFT{}).Where("id = ?", nftID).
			UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
	})
}

func (r *GormNFTRepository) IncrementNFTCount(nftID uint, countType string) error {
	result := r.db.Model(&models.NFT{}).Where("id = ?", nftID)

	switch countType {
	case "view":
		result = result.UpdateColumn("view_count", gorm.Expr("view_count + ?", 1))
	case "transaction":
		result = result.UpdateColumn("transaction_count", gorm.Expr("transaction_count + ?", 1))
	default:
		return fmt.Errorf("无效的计数类型: %s", countType)
	}

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的NFT", nftID)
	}

	return nil
}

func (r *GormNFTRepository) GetContentFeatures(batchSize int) ([][]float32, error) {
	var allMetadatatFeatures [][]float32
	var lastID uint = 0
	for {
		var batch []models.NFT

		var count int64
		if err := r.db.Model(&models.NFT{}).Count(&count).Error; err != nil {
			return nil, err
		}
		if count == 0 {
			return nil, nil // 表长度为0，跳过处理
		}

		if err := r.db.Select("id, ContentFeature").
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

func (r *GormNFTRepository) CategorizationTask() error {
	// Categorize Hot NFTs (Top 100 by ViewCount)
	if err := categorizeHotNFTs(r.db); err != nil {
		return fmt.Errorf("error categorizing hot NFTs: %w", err)
	}

	// Categorize Bestseller NFTs (Top 100 by TransactionCount)
	if err := categorizeBestsellerNFTs(r.db); err != nil {
		return fmt.Errorf("error categorizing bestseller NFTs: %w", err)
	}

	// Categorize Newest NFTs (Top 100 by CreateTime)
	if err := categorizeNewestNFTs(r.db); err != nil {
		return fmt.Errorf("error categorizing newest NFTs: %w", err)
	}

	return nil
}

func categorizeHotNFTs(db *gorm.DB) error {
	var nfts []models.NFT
	if err := db.Order("view_count desc").Limit(100).Find(&nfts).Error; err != nil {
		return err
	}

	for _, nft := range nfts {
		if err := updateNFTCategory(db, nft.ID, models.CategoryHot); err != nil {
			return err
		}
	}
	return nil
}

func categorizeBestsellerNFTs(db *gorm.DB) error {
	var nfts []models.NFT
	if err := db.Order("transaction_count desc").Limit(100).Find(&nfts).Error; err != nil {
		return err
	}

	for _, nft := range nfts {
		if err := updateNFTCategory(db, nft.ID, models.CategoryBestseller); err != nil {
			return err
		}
	}
	return nil
}

func categorizeNewestNFTs(db *gorm.DB) error {
	var nfts []models.NFT
	if err := db.Order("created_at desc").Limit(100).Find(&nfts).Error; err != nil {
		return err
	}

	for _, nft := range nfts {
		if err := updateNFTCategory(db, nft.ID, models.CategoryNewest); err != nil {
			return err
		}
	}
	return nil
}

func updateNFTCategory(db *gorm.DB, nftID uint, category models.NFTCategory) error {
	return db.Model(&models.NFT{}).Where("id = ?", nftID).
		UpdateColumn("categories", gorm.Expr("ARRAY(SELECT DISTINCT UNNEST(ARRAY_APPEND(categories, ?)))", category)).Error
}
