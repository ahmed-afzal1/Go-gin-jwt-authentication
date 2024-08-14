package repositories

import (
	"github.com/ahmed-afzal1/go-auth/config"
	"github.com/ahmed-afzal1/go-auth/helper"
	"github.com/ahmed-afzal1/go-auth/models"
)

func SignUp(user models.User) (models.User, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return user, err
	}

	token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, user.ID)
	user.Token = &token
	user.Refresh_token = &refreshToken

	return user, nil
}
