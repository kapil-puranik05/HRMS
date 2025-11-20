package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		c.Set("email", email)
		c.Next()
	}
}
