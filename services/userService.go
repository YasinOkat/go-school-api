package services

import (
	"errors"

	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUsernameExists = errors.New("username already exists")

func CreateUser(userCreate models.UserCreate) error {

	existingUser, err := repositories.GetUserByUsername(userCreate.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return ErrUsernameExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.UserCreate{
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

func GetUsers(filterByActive bool) ([]models.User, error) {
	return repositories.GetUsers(filterByActive)
}

func DeleteUser(userID uint) error {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}

	return repositories.DeleteUser(userID)
}

func GetUserByID(userID uint) (*models.User, error) {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
