package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"
)

type NFTHandler struct {
	nftService  *service.NFTService
	ipfsService *service.IPFSService
}

func NewNFTHandler(nftService *service.NFTService, ipfsService *service.IPFSService) *NFTHandler {
	return &NFTHandler{
		nftService:  nftService,
		ipfsService: ipfsService,
	}
}

// Check for plagiarism and store in IPFS
func (h *NFTHandler) GetFeatures(c *gin.Context) {
	var req models.Metadata

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	summaryFeature, err := utils.GetFeatures([]string{req.Summary})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT, get feature failed"})
		return
	}
	contentFeature, err := utils.GetFeatures([]string{req.Content})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT, get feature failed"})
		return
	}

	var batchSize int = 1000
	var confThread float32 = 0.8
	similarityHigh, err := h.nftService.CheckSimilarity(contentFeature[0], confThread, batchSize)
	if err != nil {
		log.Fatal("calculate similarity failed")
	}
	if similarityHigh {
		c.JSON(http.StatusBadRequest, gin.H{"error": "similarityHigh"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"summaryFeature": summaryFeature, "contentFeature": contentFeature})
}

// CreateNFT handles the creation of a new NFT
func (h *NFTHandler) CreateNFT(c *gin.Context) {
	var req struct {
		TokenID         string       `json:"tokenId" binding:"required"`
		ContractAddress string       `json:"contractAddress" binding:"required"`
		MetadataURI     string       `json:"metadataUri" binding:"required"`
		SummaryFeature [512]float32 `json:"summaryFeature" binding:"required"`
		ContentFeature [512]float32 `json:"metadata" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the authenticated context
	userID, _ := c.Get("userID")
	creatorID := userID.(uint)
	nftID, err := h.nftService.CreateNFT(req.TokenID, req.ContractAddress, creatorID, creatorID, req.MetadataURI, req.SummaryFeature, req.ContentFeature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": nftID})
}

// get NFT info by id
func (h *NFTHandler) GetNFTByID (c *gin.Context) {
	var req struct {
		NFTId uint `json:"nftId" binging:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nft, err := h.nftService.GetNFTDetails(req.NFTId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to select nft details"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"nft": nft})
}

func (h *NFTHandler) GetNFTsByCreator(c *gin.Context) {
	var req struct {
		Creator uint `json:"creator" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	nfts, err := h.nftService.ListNFTByCreator(req.Creator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NFTs"})
		return
	}

	if len(nfts) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this creator", "data": []models.NFT{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    nfts,
		"count":   len(nfts),
	})
}

func (h *NFTHandler) GetNFTByClassification(c *gin.Context) {
	var req struct {
		Classification string `json:"classification"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	nfts, err := h.nftService.GetNFTByClassification(req.Classification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NFTs"})
		return
	}

	if len(nfts) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this classification", "data": []models.NFT{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    nfts,
		"count":   len(nfts),
	})
}

func (h *NFTHandler) GetNFTBySummary(c *gin.Context) {
	var req struct {
		Summary	string `json:"summary"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if len(req.Summary) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input can not be empty"})
		return
	}

	feature, err := utils.GetFeatures([]string{req.Summary})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get feature failed"})
		return
	}

	nfts, err := h.nftService.GetNFTByFeature(feature)
}

