package service

import (
	"errors"
	"log"
	"runtime"
	"sort"
	"sync"
	"time"

	"nftPlantform/models"
	"nftPlantform/models/dto"
	"nftPlantform/repository"
	"nftPlantform/utils"
)

type NFTService struct {
	nftRepo *repository.GormNFTRepository
}

func NewNFTService(nftRepo *repository.GormNFTRepository) *NFTService {
	return &NFTService{
		nftRepo: nftRepo,
	}
}

// CreateNFT 创建新的NFT
func (s *NFTService) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, summaryFeature, metadataFeature [512]float32, grade, subject, topic string, price float64) (uint, error) {
	grade_, err := s.nftRepo.FindOrCreateGrade(grade)
	if err != nil {
		return 0, err
	}

	subject_, err := s.nftRepo.FindOrCreateSubject(subject)
	if err != nil {
		return 0, err
	}

	topic_, err := s.nftRepo.FindOrCreateTopic(topic)
	if err != nil {
		return 0, err
	}
	return s.nftRepo.CreateNFT(tokenID, contractAddress, ownerID, creatorID, metadataURI, summaryFeature, metadataFeature, grade_, subject_, topic_, price)
}

// GetNFTDetails 获取NFT详情
func (s *NFTService) GetNFTDetails(id uint) (*models.NFT, error) {
	err := s.nftRepo.IncrementNFTCount(id, "view")
	if err != nil {
		return nil, err
	}
	return s.nftRepo.GetNFTByID(id)
}

func (s *NFTService) SetNFTUnSaleByID(nftID uint) error {
	return s.nftRepo.SetNFTUnSaleByID(nftID)
}

// 点赞
func (s *NFTService) LikeNFTByID(userID, nftID uint) error {
	return s.nftRepo.LikeNFTByID(userID, nftID)
}

// 通过NFT的详细信息查询
func (s *NFTService) GetNFTByDetails(query dto.NFTQuery) ([]*models.NFT, error) {
	nfts, err := s.nftRepo.GetNFTByDetails(query)
	if err != nil {
		return nil, err
	}
	var res []*models.NFT
	if query.Keyword != nil && len(*query.Keyword) != 0 {
		var nftsWithSimilarity []*models.NFTWithSimilarity
		targetFeature, err := utils.GetFeatures([]string{*query.Keyword})
		if err != nil {
			return nil, err
		}
		for _, nft := range nfts {
			// 将 []byte 转换为 []float32
			summaryFeature, err := utils.BlobToFloat32Array(nft.SummaryFeature)
			if err != nil {
				return nil, err
			}
			// 计算相似度（这里使用余弦相似度作为示例）
			similarity := utils.CalculateSimilarity(targetFeature, &summaryFeature)

			nftsWithSimilarity = append(nftsWithSimilarity, &models.NFTWithSimilarity{
				NFT:        nft,
				Similarity: similarity,
			})
		}

		// 根据相似度排序
		sort.Slice(nftsWithSimilarity, func(i, j int) bool {
			return nftsWithSimilarity[i].Similarity > nftsWithSimilarity[j].Similarity
		})
		for _, value := range nftsWithSimilarity {
			res = append(res, value.NFT)
		}
	} else {
		res = nfts
	}
	// 实现分页逻辑
	totalNFTs := len(res)
	pageSize := int(*query.PageSize)
	page := int(*query.Page)

	// 计算起始和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	// 检查索引是否越界
	if startIndex >= totalNFTs {
		return []*models.NFT{}, nil // 返回空切片
	}
	if endIndex > totalNFTs {
		endIndex = totalNFTs
	}

	// 返回分页后的结果
	return res[startIndex:endIndex], nil
}

// 根据分类查询NFT
func (s *NFTService) GetNFTByClassification(classification string) ([]*models.NFT, error) {
	return s.nftRepo.GetNFTByClassification(classification)
}

// 根据年级查学科
func (s *NFTService) GetSubjectByGrade(grade uint) (*[]models.Subject, error) {
	return s.nftRepo.GetSubjectByGrade(grade)
}

// get latest number of NFT
func (s *NFTService) GetLatestNFT(number uint) (*[]models.NFT, error) {
	return s.nftRepo.GetLatestNFT(number)
}

func (s *NFTService) GetHottestNFT(number uint) (*[]models.NFT, error) {
	return s.nftRepo.GetHottestNFT(number)
}

func (s *NFTService) GetHighTradingNFT(number uint) (*[]models.NFT, error) {
	return s.nftRepo.GetHighTradingNFT(number)
}

func (s *NFTService) GetFavoriteTopic(userID uint) (*models.Topic, error) {
	return s.nftRepo.GetFavoriteTopic(userID)
}

// ListNFTsByOwner 列出用户拥有的所有NFT
func (s *NFTService) ListNFTsByOwner(ownerID uint) ([]*models.NFT, error) {
	return s.nftRepo.GetNFTsByOwnerID(ownerID)
}

// ListNFTByCreator 列出所有特定作者的NFT
func (s *NFTService) ListNFTByCreator(creatorID uint) ([]*models.NFT, error) {
	return s.nftRepo.GetNFTsByCreatorID(creatorID)
}

func (s *NFTService) GetNFTByTopicAndType(topicId *uint, typeId *uint, limit uint) (*[]models.NFT, error) {
	return s.nftRepo.GetNFTByTopicAndType(topicId, typeId, limit)
}

func (s *NFTService) GetTopicBySubjectAndGrade(subjectId, gradeId []*uint) ([]dto.IDName, error) {
	return s.nftRepo.GetTopicBySubjectAndGrade(subjectId, gradeId)
}

func (s *NFTService) GetGrade() (*[]dto.IDName, error) {
	return s.nftRepo.GetGrade()
}

// TransferNFT 转移NFT所有权
func (s *NFTService) TransferNFT(nftID, fromUserID, toUserID uint) error {
	nft, err := s.nftRepo.GetNFTByID(nftID)
	if err != nil {
		return err
	}

	if nft.OwnerID != fromUserID {
		return errors.New("user does not own this NFT")
	}

	nft.OwnerID = toUserID
	return s.nftRepo.UpdateNFT(nft)
}

// UpdateNFTMetadata 更新NFT元数据
func (s *NFTService) UpdateNFTMetadata(nftID uint, newMetadataURI string) error {
	nft, err := s.nftRepo.GetNFTByID(nftID)
	if err != nil {
		return err
	}

	nft.MetadataURI = newMetadataURI
	return s.nftRepo.UpdateNFT(nft)
}

// UpdateNFTCategory 启动一个goroutine来定期更新NFT分类
func (s *NFTService) UpdateNFTCategory(interval int) error {
	go func(interval int) {
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := s.nftRepo.CategorizationTask(); err != nil {
					log.Printf("Error during categorization task: %v", err)
				}
			}
		}
	}(interval)

	return nil
}

// 计算NFT的相似度，超过阈值的返回true，否则返回false
func (s *NFTService) CheckSimilarity(feature *[]float32, threshold float32, batchsize int) (bool, error) {
	allFeatures, err := s.nftRepo.GetContentFeatures(batchsize)
	if err != nil {
		return false, err
	}

	if len(allFeatures) == 0 {
		return false, nil
	}

	numCPU := runtime.NumCPU()
	if len(allFeatures) < numCPU*4 {
		// 对于小数据集，使用单线程处理
		for _, vec := range allFeatures {
			if utils.CalculateSimilarity(&vec, feature) > threshold {
				return true, nil
			}
		}
		return false, nil
	}

	chunkSize := (len(allFeatures) + numCPU - 1) / numCPU
	var wg sync.WaitGroup
	resultChan := make(chan bool, numCPU)

	for i := 0; i < numCPU; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(allFeatures) {
			end = len(allFeatures)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for _, vec := range allFeatures[start:end] {
				if utils.CalculateSimilarity(&vec, feature) > threshold {
					select {
					case resultChan <- true:
					default:
						// 如果通道已满，我们不等待
					}
					return
				}
			}
		}(start, end)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for range resultChan {
		return true, nil
	}

	return false, nil
}
