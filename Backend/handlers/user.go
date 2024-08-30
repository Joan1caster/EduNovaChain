package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nftPlantform/models"
	"nftPlantform/service"
	"nftPlantform/utils"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
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
