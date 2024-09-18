package service

import (
	"context"
	"errors"
	"time"

	"nftPlantform/config"
	"nftPlantform/internal/database"
	"nftPlantform/models"
	"nftPlantform/repository"
	"nftPlantform/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"github.com/spruceid/siwe-go"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repository.GormUserRepository
}

func NewUserService(userRepo *repository.GormUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

var jwtSecret = []byte(config.AppConfig.JwtSecret)

func (s *UserService) AuthenticateUser(address, signature, message string) (*models.User, error) {
	// 验证签名
	if !utils.VerifySignature(address, signature, message) {
		logrus.Warnf("Invalid signature for address %s", address)
		return nil, errors.New("invalid signature")
	}

	// 检查用户是否存在，如果不存在则创建新用户
	user, err := s.userRepo.GetUserByWalletAddress(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Infof("User not found for address %s, creating new user", address)
			userID, err := s.userRepo.CreateUser("", "", "", address)
			if err != nil {
				logrus.Errorf("Failed to create user for address %s: %v", address, err)
				return nil, err
			}
			user, err = s.userRepo.GetUserByID(userID)
			if err != nil {
				logrus.Errorf("Failed to retrieve newly created user for address %s: %v", address, err)
				return nil, err
			}
		} else {
			logrus.Errorf("Error retrieving user for address %s: %v", address, err)
			return nil, err
		}
	}

	logrus.Infof("User %s authenticated successfully", address)
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
		"chainId":   "11155111",
	}
	message, err := siwe.InitMessage(domain, walletAddress, uri, nonce, options)
	if err != nil {
		logrus.Errorf("Failed to generate SIWE message for wallet %s: %v", walletAddress, err)
		return "", err
	}

	redisClient := database.GetRedis()
	err = redisClient.Set(context.Background(), "nonce:"+walletAddress, nonce, 60*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Failed to store nonce for wallet %s: %v", walletAddress, err)
		return "", err
	}

	logrus.Infof("Generated SIWE message for wallet %s", walletAddress)
	return message.String(), nil
}

// 根据钱包查地址
func (s *UserService) GetUserByWallet(walletAddress string) (*models.User, error) {
	user, err := s.userRepo.GetUserByWalletAddress(walletAddress)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

// Login 登录
func (s *UserService) Login(messageStr, signature string) (*models.User, string, error) {
	// 解析消息
	parsedMessage, err := siwe.ParseMessage(messageStr)
	if err != nil {
		logrus.Errorf("Failed to parse SIWE message: %v", err)
		return nil, "", err
	}
	wallet := parsedMessage.GetAddress().Hex()

	// 从 Redis 中获取 nonce 并验证
	redisClient := database.GetRedis()
	storedNonce, err := redisClient.Get(context.Background(), "nonce:"+wallet).Result()
	if err != nil || storedNonce != parsedMessage.GetNonce() {
		logrus.Warnf("Invalid nonce or player not found for wallet %s", wallet)
		return nil, "", errors.New("invalid nonce or player not found")
	}

	// 验证签名
	if _, err := parsedMessage.VerifyEIP191(signature); err != nil {
		logrus.Errorf("Failed to verify signature for wallet %s: %v", wallet, err)
		return nil, "", err
	}

	// 签名验证通过后，清空 nonce
	err = redisClient.Del(context.Background(), "nonce:"+wallet).Err()
	if err != nil {
		logrus.Errorf("Failed to delete nonce for wallet %s: %v", wallet, err)
		return nil, "", err
	}

	// 从数据库中查询用户，如果不存在则创建新用户
	user, err := s.userRepo.GetUserByWalletAddress(wallet)
	if err != nil {
		logrus.Errorf("Error retrieving user for wallet %s: %v", wallet, err)
		return nil, "", err
	}

	// 生成新的 UUID 并存储到 Redis，覆盖之前的会话
	newUUID := utils.GenerateNonce()
	err = redisClient.Set(context.Background(), "uuid:"+wallet, newUUID, 12*time.Hour).Err()
	if err != nil {
		logrus.Errorf("Failed to store UUID for wallet %s: %v", wallet, err)
		return nil, "", err
	}

	// 生成 JWT
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &models.Claims{
		UserID: user.ID,
		Wallet: wallet,
		UUID:   newUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		logrus.Errorf("Failed to generate JWT for wallet %s: %v", wallet, err)
		return nil, "", err
	}
	logrus.Infof("User %s logged in successfully", wallet)
	return user, tokenString, nil
}

func (s *UserService) GetUserMostVisitedTopic(userID uint) (*models.Topic, error) {
	return s.userRepo.GetUserMostVisitedTopic(userID)
}
