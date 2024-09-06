package handlers

import (
	"log"
	"net/http"
	"strconv"

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
		SummaryFeature  [512]float32 `json:"summaryFeature" binding:"required"`
		ContentFeature  [512]float32 `json:"metadata" binding:"required"`
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
func (h *NFTHandler) GetNFTByID(c *gin.Context) {
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

func (h *NFTHandler) GetLatestNFT(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number parameter"})
		return
	}
	nfts, err := h.nftService.GetLatestNFT(uint(number))
	if err != nil || len(*nfts) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NFTs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    nfts,
		"count":   len(*nfts),
	})
}

func (h *NFTHandler) GetNFTByTopicAndType(c *gin.Context) {
	var req struct {
		TopicId *uint `json:"topicId"`
		TypeId  *uint `json:"typeId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	limit := uint(10)
	nfts, err := h.nftService.GetNFTByTopicAndType(req.TopicId, req.TypeId, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get NFT By Topic And Type from database error"})
	}
	if len(*nfts) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this classification", "data": []models.NFT{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    nfts,
		"count":   len(*nfts),
	})
}

func (h *NFTHandler) GetGradeList(c *gin.Context) {
	grades, err := h.nftService.GetGrade()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get grade from database error"})
		return
	}
	if len(*grades) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No grades found from database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "grades retrieved successfully",
		"data":    grades,
		"count":   len(*grades),
	})
}

func (h *NFTHandler) GetSubjectByGrade(c *gin.Context) {
	gradeStr := c.Param("grade")
	gradeId, err := strconv.Atoi(gradeStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "input data form error", "data": gradeStr})
	}
	subjects, err := h.nftService.GetSubjectByGrade(uint(gradeId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GetSubjectByGrade from database error"})
	}
	if len(*subjects) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this classification", "data": []models.NFT{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    subjects,
		"count":   len(*subjects),
	})
}

func (h *NFTHandler) GetTopicBySubjectAndGrade(c *gin.Context) {
	var req struct {
		SubjectId *uint `json: subjectId`
		GradeId *uint `json: gradeId`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	topics, err := h.nftService.GetTopicBySubjectAndGrade(req.SubjectId, req.GradeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GetTopicBySubjectAndGrade from database error"})
	}
	if len(*topics) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this classification", "data": []models.NFT{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    topics,
		"count":   len(*topics),
	})

}

func (h *NFTHandler) GetNFTBySummary(c *gin.Context) {
	var req struct {
		Summary string `json:"summary"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	if len(req.Summary) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input can not be empty"})
		return
	}

	_, err := utils.GetFeatures([]string{req.Summary})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get feature failed"})
		return
	}

	// nfts, err := h.nftService.GetNFTByFeature(feature)
}
