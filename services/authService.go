package services

import (
	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/repositories"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByUsername(username string) (*models.User, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
