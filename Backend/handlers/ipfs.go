package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"
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
	var nftData models.NFTData
	if err := c.ShouldBindJSON(&nftData); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	hash, err := i.ipfsService.UploadData(nftData)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	// TODO: 存储 hash 到数据库
	utils.Success(c, gin.H{"hash": hash})
}

func (i *IPFSHandler) GetData(c *gin.Context) {
	hash := c.Query("hash")
	nftData, err := i.ipfsService.GetData(hash)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, nftData)
}
