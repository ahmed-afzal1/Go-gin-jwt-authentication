package middleware

import (
	"net/http"

	"github.com/ahmed-afzal1/go-auth/helper"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("token")

	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization header provided"})
		c.Abort()
		return
	}

	claims, err := helper.ValidateToken(token)

	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Set("email", claims.Email)
	c.Set("first_name", claims.First_name)
	c.Set("last_name", claims.Last_name)
	c.Set("uid", claims.Uid)
	c.Set("user_type", claims.User_type)

	c.Next()
}
