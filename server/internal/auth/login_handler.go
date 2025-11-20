package auth

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if body.Email != os.Getenv("TEST_EMAIL") || body.Password != os.Getenv("TEST_PASSWORD") {
		c.JSON(401, gin.H{"error": "invalid email/password"})
		return
	}
	token, err := GenerateToken(body.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
