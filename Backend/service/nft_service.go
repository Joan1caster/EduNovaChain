package service

import (
	"errors"
	"runtime"
	"sync"

	"nftPlantform/api"
	"nftPlantform/models"
	"nftPlantform/utils"
)

type NFTService struct {
	nftRepo api.NFTRepository
}

func NewNFTService(nftRepo api.NFTRepository) *NFTService {
	return &NFTService{
		nftRepo: nftRepo,
	}
}

// CreateNFT 创建新的NFT
func (s *NFTService) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string, summaryFeature, metadataFeature [512]float32) (uint, error) {
	return s.nftRepo.CreateNFT(tokenID, contractAddress, ownerID, creatorID, metadataURI, summaryFeature, metadataFeature)
}

// GetNFTDetails 获取NFT详情
func (s *NFTService) GetNFTDetails(id uint) (*models.NFT, error) {
	return s.nftRepo.GetNFTByID(id)
}

// ListNFTsByOwner 列出用户拥有的所有NFT
func (s *NFTService) ListNFTsByOwner(ownerID uint) ([]*models.NFT, error) {
	return s.nftRepo.GetNFTsByOwnerID(ownerID)
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

// 计算NFT的相似度，超过阈值的返回true，否则返回false
func (s *NFTService) CheckSimilarity(feature [512]float32, threshold float32, batchsize int) (bool, error) {
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
			if utils.CalculateSimilarity(vec, feature) > threshold {
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
				if utils.CalculateSimilarity(vec, feature) > threshold {
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
