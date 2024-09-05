package handlers

import (
	"net/http"
	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetFavoriteTopic(c *gin.Context) {
	userID, _ := c.Get("userID")
	topic, err := h.userService.GetUserMostVisitedTopic(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create NFT"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"topic": topic})
}

func (u *UserHandler) GetSIWEMessage(c *gin.Context) {
	wallet := c.Query("wallet")
	message, err := u.userService.GenerateSIWEMessage(wallet)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
	}
	utils.Success(c, message)
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
