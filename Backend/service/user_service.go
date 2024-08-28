package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spruceid/siwe-go"
	"nftPlantform/api"
	"nftPlantform/internal/database"
	"nftPlantform/models"
	"nftPlantform/utils"
	"time"
)

type UserService struct {
	userRepo api.UserRepository
}

var jwtSecret = []byte("c2f7e3b6f88c4e1b8d9252f3d7c74eae63f2b41e519f2e9a75c7685c4b7e89a2")

func (s *UserService) AuthenticateUser(address, signature, message string) (*models.User, error) {
	// Verify the signature
	if !utils.VerifySignature(address, signature, message) {
		return nil, errors.New("invalid signature")
	}

	// Check if the user exists, if not, create a new user
	user, err := s.userRepo.GetUserByWalletAddress(address)
	if err != nil {
		// User doesn't exist, create a new one
		userID, err := s.userRepo.CreateUser("", "", "", address)
		if err != nil {
			return nil, err
		}
		user, err = s.userRepo.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

// GenerateSIWEMessage 生成SIWE信息
func (s *UserService) GenerateSIWEMessage(walletAddress string) (string, error) {
	domain := "nftPlantform"
	uri := "https://nftPlantform.xyz"
	nonce := utils.GenerateNonce()
	statement := "Welcome to nftPlantform!"
	options := map[string]interface{}{
		"statement": statement,
		"chainId":   "1",
	}
	message, err := siwe.InitMessage(domain, walletAddress, uri, nonce, options)
	if err != nil {
		return "", err
	}
	redisClient := database.GetRedis()
	err = redisClient.Set(context.Background(), "nonce:"+walletAddress, nonce, 5*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return message.String(), nil
}

// Login 登录
func (s *UserService) Login(messageStr, signature string) (string, error) {
	// 解析消息
	parsedMessage, err := siwe.ParseMessage(messageStr)
	if err != nil {
		return "", err
	}
	wallet := parsedMessage.GetAddress().Hex()

	// 从 Redis 中获取 nonce 并验证
	redisClient := database.GetRedis()
	storedNonce, err := redisClient.Get(context.Background(), "nonce:"+wallet).Result()
	if err != nil || storedNonce != parsedMessage.GetNonce() {
		return "", errors.New("invalid nonce or player not found")
	}

	// 验证签名
	if _, err := parsedMessage.VerifyEIP191(signature); err != nil {
		return "", err
	}

	// 签名验证通过后，清空 nonce
	err = redisClient.Del(context.Background(), "nonce:"+wallet).Err()
	if err != nil {
		return "", err
	}

	// 生成新的 UUID 并存储到 Redis，覆盖之前的会话
	newUUID := utils.GenerateNonce()
	err = redisClient.Set(context.Background(), "uuid:"+wallet, newUUID, 12*time.Hour).Err()
	if err != nil {
		return "", err
	}

	// 生成 JWT
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &models.Claims{
		Wallet: wallet,
		UUID:   newUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
