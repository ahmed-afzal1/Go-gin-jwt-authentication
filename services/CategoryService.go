package services

import (
	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/repositories"
)

func GetAllCategories() ([]models.Category, error) {
	return repositories.GetAllCategories()
}

func CreateCategory(category models.Category) (models.Category, error) {
	return repositories.CreateCategory(category)
}
