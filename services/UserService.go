package services

import (
	"log"

	"github.com/ahmed-afzal1/go-auth/models"
	"github.com/ahmed-afzal1/go-auth/repositories"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func SignUp(user models.User) (models.User, error) {
	password := HashPassword(*user.Password)
	user.Password = &password

	savedUser, err := repositories.SignUp(user)

	if err != nil {
		return models.User{}, err
	}

	return savedUser, nil
}
