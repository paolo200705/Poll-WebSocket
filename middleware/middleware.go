package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Otavio-Fina/live-websocket/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func LoginHandler(c *gin.Context) {
	userID := uuid.New().String()

	claims := &models.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30000 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(models.JwtSecret)

	if err != nil {
		fmt.Printf("erro ao gerar o token: %v\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o Token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"token":   tokenString,
	})
}

func AuthMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")

		if tokenString == "" {
			tokenString = c.Query("token")
		}

		// Remove "Bearer " se estiver presente
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Invalido paizao"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
