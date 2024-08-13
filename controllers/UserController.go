package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
