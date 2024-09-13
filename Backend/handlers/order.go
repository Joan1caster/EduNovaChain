package handlers

import (
	"io"
	"net/http"
	"time"

	"nftPlantform/common"
	"nftPlantform/models"
	"nftPlantform/models/dto"
	"nftPlantform/service"
	"nftPlantform/utils"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
	nftService   *service.NFTService
	tradeService *service.NFTTrade
}

func NewOrderHandler(orderService *service.OrderService, nftService *service.NFTService, tradeService *service.NFTTrade) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		nftService:   nftService,
		tradeService: tradeService,
	}
}

// LatestNFT godoc
// @Description latest NFT list for sale
// @Tags nft
// @Accept Null
// @Produce json
// @Param Null
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nft/list [post]
func (h *OrderHandler) LatestNFT(c *gin.Context) {

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

	wallet, exists := c.Get("wallet")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var newOrderReq dto.OrderMessage
	if err := c.BindJSON(&newOrderReq); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// 验证NFT所有权
	nft, err := h.nftService.GetNFTDetails(newOrderReq.NFTId)

	if nft == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "trade NFT do not exist"})
		return
	}

	if err != nil || nft.Owner.WalletAddress != wallet.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't own this NFT"})
		return
	}

	orderID, err := h.orderService.ListNFTForSale(nft.Owner.ID, newOrderReq.NFTId, newOrderReq.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	UserID, exists := c.Get("userID")
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

	err := h.orderService.ValidateOrderStatus(req.OrderID, UserID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.orderService.CancelOrder(req.OrderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	orders, err := h.orderService.GetCompletedOrdersByNFTID(req.nftId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select history orders"})
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No NFTs found for this creator", "data": []models.Order{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": "orders retrieved successfully",
		"data":   orders,
		"count":  len(orders),
	})
}

func (h *OrderHandler) BuyNFT(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req BuyNFTRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	order, err := h.orderService.GetOrderByID(req.OrderID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.orderService.ValidateOrderStatus(order.ID, order.SellerID); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Step 2: 创建交易
	//err = h.tradeService.CreateTransaction(order.ID, order.NFTID, userID.(uint), req.TxHash, fmt.Sprintf("%f",order.Price))
	//if err != nil {
	//	utils.Error(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	if _, exists := common.TxStatusChannels.Get(req.TxHash); !exists {
		common.NewTxStatusChannel(req.TxHash)
	} else {
		utils.Error(c, http.StatusInternalServerError, "tx has been listened, please run <orders/status:txHash> to check")
		return
	}
	go h.tradeService.StartTransactionListener(order.NFTID, order.ID, userID.(uint), req.TxHash)

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "submitted",
		"txHash":  req.TxHash,
		"message": "Transaction submitted, awaiting confirmation",
	})
}

func (h *OrderHandler) TransactionStatus(c *gin.Context) {
	txHash := c.Param("txHash")

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	statusCh, exists := common.TxStatusChannels.Get(txHash)
	if !exists || statusCh == nil {
		utils.Error(c, http.StatusInternalServerError, "get status failed, channel do not exists")
	}
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-statusCh:
			if !ok {
				c.SSEvent("message", gin.H{"status": "maybe tradition has been confirmed", "txHash": txHash})
				return false
			}
			c.SSEvent("message", gin.H{"status": msg, "txHash": txHash})
			statusCh <- msg
			return false // 继续流
		case <-time.After(5 * time.Second):
			// 超时处理
			c.SSEvent("wa", "Timeout waiting for status update")
			return false
		}
	})
}

type BuyNFTRequest struct {
	OrderID uint   `json:"order_id" binding:"required"`
	TxHash  string `json:"tx_hash" binding:"required"`
}
