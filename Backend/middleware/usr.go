package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"nftPlantform/internal/database"
	"nftPlantform/models"
)

var jwtSecret = []byte("c2f7e3b6f88c4e1b8d9252f3d7c74eae63f2b41e519f2e9a75c7685c4b7e89a2")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		redisClient := database.GetRedis()
		storedUUID, err := redisClient.Get(context.Background(), "uuid:"+claims.Wallet).Result()
		if errors.Is(err, redis.Nil) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session not found"})
			c.Abort()
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			c.Abort()
			return
		}
		if storedUUID != claims.UUID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired or invalid"})
			c.Abort()
			return
		}
		c.Set("wallet", claims.Wallet)
		c.Next()
	}
}
