package controllers

import (
	"net/http"
	"strconv"

	"github.com/YasinOkat/go-school-api/models"
	"github.com/YasinOkat/go-school-api/services"
	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new user with the input payload
// @Tags students
// @Accept  json
// @Produce  json
// @Param user body models.StudentUserCreate true "Create Student"
// @Success 201 {object} models.StudentUserCreate
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse "username already exists"
// @Failure 500 {object} models.ErrorResponse
// @Router /students [post]
func CreateStudent(c *gin.Context) {
	var studentUser models.StudentUserCreate
	if err := c.ShouldBindJSON(&studentUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.UserCreate{
		Username:    studentUser.Username,
		Password:    studentUser.Password,
		FirstName:   studentUser.FirstName,
		LastName:    studentUser.LastName,
		PhoneNumber: studentUser.PhoneNumber,
		Email:       studentUser.Email,
		UserTypeID:  studentUser.UserTypeID,
		Status:      studentUser.Status,
	}

	err := services.CreateUser(user)
	if err == services.ErrUsernameExists {
		c.JSON(http.StatusConflict, models.ErrorResponse{Error: "username already exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser, err := services.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	student := models.StudentCreate{
		UserID:  int(newUser.ID),
		MajorID: studentUser.MajorID,
		Status:  studentUser.Status,
	}

	err = services.CreateStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, studentUser)
}

// GetStudents godoc
// @Summary Fetch all students
// @Description Fetch All students
// @Tags students
// @Produce  json
// @Success 200 {object} []models.StudentRead
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /students/ [get]
func GetStudents(c *gin.Context) {
	students, err := services.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "some error happened"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// GetStudentCourses godoc
// @Summary Get student courses
// @Description Get student courses
// @Tags students
// @Produce  json
// @Param id path int true "Student ID"
// @Success 200 {object} []models.StudentCourse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse "student not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /students/{id}/courses [get]
func GetStudentCourses(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid student ID"})
		return
	}

	studentCourse, err := services.GetStudentCourses(studentID)

	if err == services.ErrUserNotFound {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "student not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, studentCourse)
}

// SelectCourse godoc
// @Summary Select a course for a student
// @Description Select a course for a student if the course matches the student's major
// @Tags students
// @Accept  json
// @Produce  json
// @Param studentCourse body models.StudentCourseSelect true "Select Course"
// @Success 200 {string} string "course selected successfully"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse "student does not exist"
// @Failure 409 {object} models.ErrorResponse "course does not match student's major"
// @Failure 500 {object} models.ErrorResponse
// @Router /students/courses [post]
func SelectCourse(c *gin.Context) {
	var studentCourse models.StudentCourseSelect
	if err := c.ShouldBindJSON(&studentCourse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.SelectCourse(studentCourse)
	if err == services.ErrCourseMajorMismatch {
		c.JSON(http.StatusConflict, models.ErrorResponse{Error: "course does not match student's major"})
		return
	}
	if err == services.ErrStudentDoesNotExist {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "student does not exist"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "course selected successfully")
}
