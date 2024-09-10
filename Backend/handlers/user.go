package handlers

import (
	"net/http"
	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
	nftService  *service.NFTService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetFavoriteTopic(c *gin.Context) {
	wallet := c.Query("wallet")
	user, err := h.userService.GetUserByWallet(wallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exe GetUserByWallet "})
		return
	}
	topic, err := h.userService.GetUserMostVisitedTopic(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exe GetUserMostVisitedTopic "})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"topic": topic})
}

func (u *UserHandler) GetSIWEMessage(c *gin.Context) {
	wallet := c.Query("wallet")
	message, err := u.userService.GenerateSIWEMessage(wallet)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, message)
}

func (u *UserHandler) CheckAuth(c *gin.Context) {
	wallet, exists := c.Get("wallet")
	if !exists {
		utils.Error(c, http.StatusUnauthorized, "用户未登录")
		return
	}

	nftID := c.Param("nftID")
	intNftID, err := strconv.Atoi(nftID)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "无效的 NFT ID")
		return
	}

	nft, err := u.nftService.GetNFTDetails(uint(intNftID))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "获取 NFT 详情失败")
		return
	}

	if nft.Owner.WalletAddress != wallet {
		utils.Error(c, http.StatusUnauthorized, "无权限访问此 NFT")
		return
	}

	utils.Success(c, gin.H{"authority": true})
}

func (u *UserHandler) Login(c *gin.Context) {
	var loginReq struct {
		SignMessage string `json:"signMessage"`
		Signature   string `json:"signature"`
	}
	if err := c.BindJSON(&loginReq); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, token, err := u.userService.Login(loginReq.SignMessage, loginReq.Signature)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 合并 user 和 token 到一个响应结构体中
	type LoginResponse struct {
		User  *models.User `json:"user"`
		Token string       `json:"token"`
	}

	response := LoginResponse{
		User:  user,
		Token: token,
	}

	utils.Success(c, response)
}
