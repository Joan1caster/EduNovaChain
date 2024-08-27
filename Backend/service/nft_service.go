package service

import (
	"errors"
	"nftPlantform/api"
	"nftPlantform/models"
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
func (s *NFTService) CreateNFT(tokenID, contractAddress string, ownerID, creatorID uint, metadataURI string) (uint, error) {
	return s.nftRepo.CreateNFT(tokenID, contractAddress, ownerID, creatorID, metadataURI)
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
