package services

import (
	"errors"

	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/repositories"
)

var ErrCourseMajorMismatch = errors.New("course does not match student's major")
var ErrStudentDoesNotExist = errors.New("student does not exist")

func CreateStudent(studentCreate models.StudentCreate) error {

	student := models.StudentCreate{
		UserID:  studentCreate.UserID,
		MajorID: studentCreate.MajorID,
		Status:  studentCreate.Status,
	}

	err := repositories.CreateStudent(student)
	return err
}

func GetStudents() ([]models.StudentRead, error) {
	return repositories.GetStudents()
}

func GetStudentCourses(studentID int) ([]models.StudentCourse, error) {
	studentCourse, err := repositories.GetStudentCourses(studentID)

	if err != nil {
		return nil, err
	}
	if studentCourse == nil {
		return nil, ErrStudentDoesNotExist
	}
	return studentCourse, err
}

func SelectCourse(studentCourseSelect models.StudentCourseSelect) error {
	studentMajorID, err := repositories.GetStudentMajor(studentCourseSelect.StudentID)
	if err != nil {
		return err
	}
	courseMajorID, err := repositories.GetCourseMajor(studentCourseSelect.CourseID)
	if err != nil {
		return err
	}

	if studentMajorID == 0 {
		return ErrStudentDoesNotExist
	}

	if studentMajorID != courseMajorID {
		return ErrCourseMajorMismatch
	}

	err = repositories.SelectCourse(studentCourseSelect)
	if err != nil {
		return err
	}

	return nil
}
