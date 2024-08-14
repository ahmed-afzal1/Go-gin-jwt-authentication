package repositories

import (
	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/models"
)

func GetAllPost() ([]models.Post, error) {
	var posts []models.Post

	if err := config.DB.Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func CreatePost(post models.Post) (models.Post, error) {
	if err := config.DB.Create(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func FindPost(id uint) (models.Post, error) {
	var post models.Post

	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		return post, err
	}

	return post, nil
}
