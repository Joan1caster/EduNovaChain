package service

import (
	"nftPlantform/api"
	"nftPlantform/models"
)

type IPFSService struct {
	client api.IPFSRepository
}

func NewIPFSService(client api.IPFSRepository) *IPFSService {
	return &IPFSService{
		client: client,
	}
}

func (s *IPFSService) GetData(ipfsHash string) (*models.Metadata, error) {
	data, err := s.client.GetData(ipfsHash)
	return data, err
}

func (s *IPFSService) UploadData(createRequest interface{}) (*models.IpfsResponseData, error) {
	ipfsData, err := s.client.UploadData(createRequest)
	return ipfsData, err
}
