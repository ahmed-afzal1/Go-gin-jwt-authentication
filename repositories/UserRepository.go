package repositories

import (
	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/models"
)

func SignUp(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
