package repositories

import (
	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/utils"
)

func CreateStudent(student models.StudentCreate) error {
	query := `
	INSERT INTO student
	(user_id, major_id, status)
	VALUES 
	(?, ?, ?)
	`

	_, err := utils.DB.Query(
		query,
		student.UserID,
		student.MajorID,
		student.Status)

	return err
}
