package handlers

import (
	"encoding/json"
	"net/http"

	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"

	"github.com/gin-gonic/gin"
)

type IPFSHandler struct {
	ipfsService *service.IPFSService
}

func NewIPFSHandler(ipfsService *service.IPFSService) *IPFSHandler {
	return &IPFSHandler{
		ipfsService: ipfsService,
	}
}

func (i *IPFSHandler) UploadData(c *gin.Context) {
	var metaData models.Metadata
	if err := c.ShouldBindJSON(&metaData); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	// hash, err := i.ipfsService.UploadData(metaData)
	jsonData, err := json.Marshal(metaData)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	hash, err := utils.UploadString(string(jsonData))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, hash)
}

func (i *IPFSHandler) GetData(c *gin.Context) {
	hash := c.Param("hash")
	// metaData, err := i.ipfsService.GetData(hash)
	var metaData models.Metadata
	bytesData, err := utils.DownloadString(hash)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.Unmarshal(bytesData, &metaData)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, metaData)
}
