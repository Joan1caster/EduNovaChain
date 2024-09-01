package handlers

import (
	"log"
	"net/http"
	_ "strconv"

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

// // BuyNFT handles the purchase of an NFT
// func (h *NFTHandler) BuyNFT(c *gin.Context) {
// 	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
// 		return
// 	}

// 	// Get the buyer's address from the authenticated context
// 	buyerAddress, _ := c.Get("walletAddress")
// 	buyerAddr := buyerAddress.(string)

// 	err = h.nftService.BuyNFT(buyerAddr, uint(orderID))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "NFT purchased successfully"})
// }

// // GetNFTs handles fetching a list of NFTs
// func (h *NFTHandler) GetNFTs(c *gin.Context) {
// 	// Parse query parameters for pagination
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

// 	nfts, total, err := h.nftService.GetNFTs(page, limit)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch NFTs"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"nfts":  nfts,
// 		"total": total,
// 		"page":  page,
// 		"limit": limit,
// 	})
// }

// // CreateOrder handles the creation of a new order (listing an NFT for sale)
// func (h *NFTHandler) CreateOrder(c *gin.Context) {
// 	var req struct {
// 		NFTID uint    `json:"nftId" binding:"required"`
// 		Price float64 `json:"price" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Get the user ID from the authenticated context
// 	userID, _ := c.Get("userID")
// 	sellerID := userID.(uint)

// 	orderID, err := h.nftService.CreateOrder(sellerID, req.NFTID, req.Price)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"id": orderID})
// }

// // GetUserProfile handles fetching the user's profile
// func (h *NFTHandler) GetUserProfile(c *gin.Context) {
// 	// Get the user ID from the authenticated context
// 	userID, _ := c.Get("userID")
// 	id := userID.(uint)

// 	user, err := h.nftService.GetUserByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }
