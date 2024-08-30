package service

import (
	"bytes"
	"encoding/json"
	"nftPlantform/api"
	"nftPlantform/models"

	"github.com/sirupsen/logrus"
)

type IPFSService struct {
	client api.IPFSRepository
}

func NewIPFSService(client api.IPFSRepository) *IPFSService {
	return &IPFSService{
		client: client,
	}
}

func (s *IPFSService) GetData(hash string) (*models.NFTData, error) {
	var nftMetaData models.NFTData
	data, err := s.client.GetData(hash)
	if err != nil {
		logrus.Errorf("Failed to get data from IPFS: %v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &nftMetaData)
	if err != nil {
		logrus.Errorf("Failed to unmarshal NFT metadata: %v", err)
		return nil, err
	}
	return &nftMetaData, nil
}

func (s *IPFSService) UploadData(nftMetaData models.NFTData) (string, error) {
	jsonData, err := json.Marshal(nftMetaData)
	if err != nil {
		logrus.Errorf("Failed to marshal NFT metadata: %v", err)
		return "", err
	}
	hash, err := s.client.UploadData(bytes.NewReader(jsonData))
	if err != nil {
		logrus.Errorf("Failed to upload data to IPFS: %v", err)
		return "", err
	}
	logrus.Infof("Successfully uploaded NFT metadata to IPFS with hash: %s", hash)
	return hash, nil
}
