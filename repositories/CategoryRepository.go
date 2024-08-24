package repositories

import (
	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	if err := config.DB.Find(&categories).Error; err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

func CreateCategory(category models.Category) (models.Category, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}
