package handlers

import (
	"net/http"

	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderService *service.OrderService
	nftService   *service.NFTService
	tradeService   *service.NFTTrade
}

func NewOrderHandler(OrderService *service.OrderService, nftService *service.NFTService, tradeService *service.NFTTrade) *OrderHandler {
	return &OrderHandler{
		OrderService: OrderService,
		nftService:   nftService,
		tradeService: tradeService,
	}
}

// ListNFT godoc
// @Summary List an NFT for sale
// @Description Create a new order to list an NFT for sale
// @Tags nft
// @Accept json
// @Produce json
// @Param nft_id query uint true "NFT ID"
// @Param price query float64 true "Listing price"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nft/list [post]
func (h *OrderHandler) ListNFT(c *gin.Context) {
	type orderMessage struct {
		nftId uint
		price float64
	}
	wallet, exists := c.Get("wallet")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var newOrderReq orderMessage
	if err := c.BindJSON(&newOrderReq); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// 验证NFT所有权
	nft, err := h.nftService.GetNFTDetails(newOrderReq.nftId)

	if err != nil || nft.Owner.WalletAddress != wallet.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't own this NFT"})
		return
	}

	orderID, err := h.OrderService.ListNFTForSale(nft.Owner.ID, newOrderReq.nftId, newOrderReq.price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "NFT listed successfully", "order_id": orderID})
}

// DelistNFT godoc
// @Summary Delist an NFT from sale
// @Description Cancel an existing order to delist an NFT from sale
// @Tags nft
// @Accept json
// @Produce json
// @Param order_id query uint true "Order ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nft/delist [post]
func (h *OrderHandler) DelistNFT(c *gin.Context) {
	wallet, exists := c.Get("wallet")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	type DelistNFTRequest struct {
		OrderID uint `json:"order_id" binding:"required"`
	}
	var req DelistNFTRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.OrderService.ValidateOrderStatus(req.OrderID, wallet.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate order"})
		return
	}
	err = h.OrderService.CancelOrder(req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "NFT delisted successfully"})
}

func (h *OrderHandler) GetHistoryByNFTId(c *gin.Context) {
	var req struct {
		nftId uint
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	orders, err := h.OrderService.GetCompletedOrdersByNFTID(req.nftId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to select history orders"})
	}

	if len(orders) == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this creator", "data": []models.Order{}})
			return
		}

	c.JSON(http.StatusOK, gin.H{
		"orders": "orders retrieved successfully",
		"data":    orders,
		"count":   len(orders),
	})
}

// func (h *OrderHandler) BuyNFT(c *gin.Context) {
// 	wallet, exists := c.Get("wallet")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	var req BuyNFTRequest
// 	if err := c.BindJSON(&req); err != nil {
// 		utils.Error(c, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	// 验证订单状态
// 	order, err := h.orderRepo.GetOrderByID(req.OrderID)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 		return
// 	}

// 	if order.Status != "OPEN" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Order is not open for purchase"})
// 		return
// 	}

// 	// 验证买家不是卖家
// 	if order.Seller.WalletAddress == wallet.(string) {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot buy your own NFT"})
// 		return
// 	}

// 	// 获取买家ID
// 	buyer, err := h.userRepo.GetUserByWallet(wallet.(string))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get buyer information"})
// 		return
// 	}

// 	// 启动交易流程
// 	err = h.nftTradeService.ExecuteTrade(c.Request.Context(), req.OrderID, buyer.ID, req.TxHash)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate purchase: " + err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "NFT purchase initiated successfully", "order_id": req.OrderID})
// }

// type BuyNFTRequest struct {
// 	OrderID uint   `json:"order_id" binding:"required"`
// 	TxHash  string `json:"tx_hash" binding:"required"`
// }
