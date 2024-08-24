package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/services"
	"github.com/gin-gonic/gin"
)

func GetAllPost(c *gin.Context) {
	posts, err := services.GetAllPost()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	Uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UserID not found"})
		return
	}

	userIDUint, ok := Uid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserID is of an unexpected type"})
		return
	}

	var post models.Post
	var category models.Category

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := config.DB.First(&category, post.CategoryID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Category not found!"})
		return
	}

	// validationErr := validate.Struct(post)
	// if validationErr != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
	// 	return
	// }

	post.UserID = userIDUint

	newPost, err := services.CreatePost(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, newPost)
}

func FindPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := services.FindPost(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, post)
}
