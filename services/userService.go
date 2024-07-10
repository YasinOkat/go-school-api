package services

import (
	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userCreate models.UserCreate) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username:    userCreate.Username,
		Password:    string(hashedPassword),
		FirstName:   userCreate.FirstName,
		LastName:    userCreate.LastName,
		PhoneNumber: userCreate.PhoneNumber,
		Email:       userCreate.Email,
		UserTypeID:  userCreate.UserTypeID,
		Status:      userCreate.Status,
	}

	err = repositories.CreateUser(user)
	return err
}
