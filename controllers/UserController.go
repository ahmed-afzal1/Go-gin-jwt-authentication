package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func SignUp(c *gin.Context) {
	var user models.User

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	newUser, err := services.SignUp(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, newUser)
}

func UserDetails(c *gin.Context) {
	var user models.User
	email, _ := c.Get("email")

	result := config.DB.Select("first_name", "last_name", "email", "phone").Where("email = ?", email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
