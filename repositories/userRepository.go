package repositories

import (
	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/utils"
)

func CreateUser(user models.User) error {
	result := utils.DB.Create(&user)
	return result.Error
}
