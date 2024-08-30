package handlers

//
//import (
//	"net/http"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//	"nftPlantform/service"
//)
//
//type Handler struct {
//	marketplaceService *service.NFTMarketplaceService
//}
//
//func NewHandler(marketplaceService *service.NFTMarketplaceService) *Handler {
//	return &Handler{
//		marketplaceService: marketplaceService,
//	}
//}
//
//// CreateNFT handles the creation of a new NFT
//func (h *Handler) CreateNFT(c *gin.Context) {
//	var req struct {
//		TokenID         string `json:"tokenId" binding:"required"`
//		ContractAddress string `json:"contractAddress" binding:"required"`
//		MetadataURI     string `json:"metadataUri" binding:"required"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Get the user ID from the authenticated context
//	userID, _ := c.Get("userID")
//	creatorID := userID.(uint)
//
//	nftID, err := h.marketplaceService.CreateNFT(req.TokenID, req.ContractAddress, creatorID, creatorID, req.MetadataURI)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"id": nftID})
//}
//
//// BuyNFT handles the purchase of an NFT
//func (h *Handler) BuyNFT(c *gin.Context) {
//	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
//		return
//	}
//
//	// Get the buyer's address from the authenticated context
//	buyerAddress, _ := c.Get("walletAddress")
//	buyerAddr := buyerAddress.(string)
//
//	err = h.marketplaceService.BuyNFT(buyerAddr, uint(orderID))
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "NFT purchased successfully"})
//}
//
//// GetNFTs handles fetching a list of NFTs
//func (h *Handler) GetNFTs(c *gin.Context) {
//	// Parse query parameters for pagination
//	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
//	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
//
//	nfts, total, err := h.marketplaceService.GetNFTs(page, limit)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch NFTs"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"nfts":  nfts,
//		"total": total,
//		"page":  page,
//		"limit": limit,
//	})
//}
//
//// CreateOrder handles the creation of a new order (listing an NFT for sale)
//func (h *Handler) CreateOrder(c *gin.Context) {
//	var req struct {
//		NFTID uint    `json:"nftId" binding:"required"`
//		Price float64 `json:"price" binding:"required"`
//	}
//
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Get the user ID from the authenticated context
//	userID, _ := c.Get("userID")
//	sellerID := userID.(uint)
//
//	orderID, err := h.marketplaceService.CreateOrder(sellerID, req.NFTID, req.Price)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
//		return
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"id": orderID})
//}
//
//// GetUserProfile handles fetching the user's profile
//func (h *Handler) GetUserProfile(c *gin.Context) {
//	// Get the user ID from the authenticated context
//	userID, _ := c.Get("userID")
//	id := userID.(uint)
//
//	user, err := h.marketplaceService.GetUserByID(id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"user": user})
//}
