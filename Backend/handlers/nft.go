package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"nftPlantform/models"
	"nftPlantform/models/dto"
	"nftPlantform/service"
	"nftPlantform/utils"
)

type NFTHandler struct {
	nftService  *service.NFTService
	ipfsService *service.IPFSService
	userService *service.UserService
}

func NewNFTHandler(nftService *service.NFTService, ipfsService *service.IPFSService, userService *service.UserService) *NFTHandler {
	return &NFTHandler{
		nftService:  nftService,
		ipfsService: ipfsService,
		userService: userService,
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
		c.Error(err)
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
	similarityHigh, err := h.nftService.CheckSimilarity(contentFeature, confThread, batchSize)
	if err != nil {
		log.Fatal("calculate similarity failed")
	}
	if similarityHigh {
		c.JSON(http.StatusBadRequest, gin.H{"error": "similarityHigh"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"summaryFeature": summaryFeature, "contentFeature": contentFeature})
}

// CreateNFT handles the creation of a new NFT
func (h *NFTHandler) CreateNFT(c *gin.Context) {
	var req dto.CreateNFT

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wallet, _ := c.Get("wallet")
	user, err := h.userService.GetUserByWallet(wallet.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to GetUserByWallet"})
		return
	}
	creatorID := user.ID

	nftID, err := h.nftService.CreateNFT(req.TokenID, req.ContractAddress, creatorID, creatorID, req.MetadataURI, req.SummaryFeature, req.ContentFeature, req.Grade, req.Subject, req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT, tokenID exist"})
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": nftID})
}

func (h *NFTHandler) LikeNFT(c *gin.Context) {
	nftID := c.Param("nftID")
	uintnftID, err := strconv.Atoi(nftID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get param: nftID"})
		return
	}
	wallet, _ := c.Get("wallet")
	user, err := h.userService.GetUserByWallet(wallet.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to GetUserByWallet"})
		return
	}

	err = h.nftService.LikeNFTByID(user.ID, uint(uintnftID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Like NFT"})
		return
	}
}

// get NFT info by id
func (h *NFTHandler) GetNFTByID(c *gin.Context) {
	nftID := c.Param("id")
	nftIDInt, err := strconv.Atoi(nftID)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "input nftid error")
	}
	nft, err := h.nftService.GetNFTDetails(uint(nftIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to select nft details"})
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"nft": nft})
}

func (h *NFTHandler) GetNFTsByCreator(c *gin.Context) {
	var req struct {
		Creator uint `json:"creatorID" binding:"required"`
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
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    nfts,
		"count":   len(*nfts),
	})
}

func (h *NFTHandler) GetHottestNFT(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number parameter"})
		return
	}
	nfts, err := h.nftService.GetHottestNFT(uint(number))
	if err != nil || len(*nfts) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NFTs"})
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "NFTs retrieved successfully",
		"data":    nfts,
		"count":   len(*nfts),
	})
}

func (h *NFTHandler) GetHighTradingNFT(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number parameter"})
		return
	}
	nfts, err := h.nftService.GetHighTradingNFT(uint(number))
	if err != nil || len(*nfts) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve NFTs"})
		c.Error(err)
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
		c.Error(err)
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
		"message": "Subject retrieved successfully",
		"data":    subjects,
		"count":   len(*subjects),
	})
}

func (h *NFTHandler) GetNFTByDetails(c *gin.Context) {
	var req dto.NFTQuery
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	nfts, err := h.nftService.GetNFTByDetails(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GetNFTByDetails from database error"})
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

func (h *NFTHandler) GetTopicBySubjectAndGrade(c *gin.Context) {
	var req dto.SubjectsAndGrades
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	topics, err := h.nftService.GetTopicBySubjectAndGrade(req.Subjectids, req.Gradeids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "GetTopicBySubjectAndGrade from database error"})
	}
	if len(topics) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this classification", "data": []dto.IDName{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "topic retrieved successfully",
		"data":    topics,
		"count":   len(topics),
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
