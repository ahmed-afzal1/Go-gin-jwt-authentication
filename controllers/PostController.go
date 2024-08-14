package controllers

import (
	"context"
	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {
	var post models.Post

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	validationErr := validate.Struct(post)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	UserID, _ := c.Get("uid")
	post.UserID = UserID

	newPost, err := services.CreatePost(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, newPost)
}
