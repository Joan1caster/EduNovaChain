package api

import (
	"nftPlantform/models"
)

type IPFSRepository interface {
	UploadData(createData interface{}) (*models.IpfsResponseData, error)
	GetData(hash string) (*models.Metadata, error)
}
