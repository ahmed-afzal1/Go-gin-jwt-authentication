package repositories

import (
	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/models"
)

func CreatePost(post models.Post) (models.Post, error) {
	if err := config.DB.Create(&post).Error; err != nil {
		return models.Post{}, err
	}

	return post, nil
}
