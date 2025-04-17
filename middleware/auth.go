package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var jwtKey = []byte("secret")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		// Проверка на Bearer
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Invalid or expired token",
				"details": err.Error(), // Это поможет отладке
			})
			c.Abort()
			return
		}

		// Добавляем user_id и is_admin в контекст
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload: user_id"})
			c.Abort()
			return
		}

		isAdmin, ok := claims["admin"].(bool)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload: admin"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(userID))
		c.Set("is_admin", isAdmin)

		c.Next()
	}
}
