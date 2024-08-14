package services

import (
	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/repositories"
)

func CreatePost(post models.Post) (models.Post, error) {
	return repositories.CreatePost(post)
}
