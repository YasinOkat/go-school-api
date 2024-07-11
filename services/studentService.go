package services

import (
	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/repositories"
)

func CreateStudent(studentCreate models.StudentCreate) error {

	student := models.StudentCreate{
		UserID:  studentCreate.UserID,
		MajorID: studentCreate.MajorID,
		Status:  studentCreate.Status,
	}

	err := repositories.CreateStudent(student)
	return err
}
