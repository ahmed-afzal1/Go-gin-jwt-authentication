package services

import (
	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/repositories"
)

func GetAllPost() ([]models.Post, error) {
	return repositories.GetAllPost()
}

func CreatePost(post models.Post) (models.Post, error) {
	return repositories.CreatePost(post)
}

func FindPost(id uint) (models.Post, error) {
	return repositories.FindPost(id)
}
