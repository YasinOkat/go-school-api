package repositories

import (
	"database/sql"

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

func GetStudents() ([]models.StudentRead, error) {
	query := `
	SELECT id, user_id, major_id, status FROM student
	`
	rows, err := utils.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.StudentRead
	for rows.Next() {
		var student models.StudentRead
		err := rows.Scan(&student.ID, &student.UserID, &student.MajorID, &student.Status)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func GetStudentMajor(studentID int) (int, error) {
	var majorID int
	query := `SELECT major_id FROM student WHERE id = ?`
	err := utils.DB.QueryRow(query, studentID).Scan(&majorID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return majorID, nil
}

func GetCourseMajor(courseID int) (int, error) {
	var majorID int
	query := `SELECT major_id FROM major_course WHERE course_id = ?`
	err := utils.DB.QueryRow(query, courseID).Scan(&majorID)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return majorID, nil
}

func AssignGrade(assignGrade models.Grade) error {
	query := "INSERT INTO student_course SET grade = ? WHERE student_id = ? AND course_id = ?"
	_, err := utils.DB.Exec(query, assignGrade.Grade, assignGrade.StudentID, assignGrade.CourseID)
	return err
}

func GetStudentCourses(studentID int) ([]models.StudentCourse, error) {
	query := `
	SELECT
		s.id AS StudentID,
		c.id AS CourseID,
		u.username AS Username,
		c.name AS CourseName
	FROM
		student s
	LEFT JOIN
		user u ON u.id = s.user_id
	LEFT JOIN
		student_course sc ON sc.student_id = s.id
	LEFT JOIN
		course c ON c.id = sc.course_id
	WHERE
		s.id = (?);
	`
	rows, err := utils.DB.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentCourse []models.StudentCourse
	for rows.Next() {
		var student models.StudentCourse
		err := rows.Scan(&student.StudentID, &student.CourseID, &student.Username, &student.CourseName)
		if err != nil {
			return nil, err
		}
		studentCourse = append(studentCourse, student)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return studentCourse, nil
}

func SelectCourse(studentCourseSelect models.StudentCourseSelect) error {
	query := `
	INSERT INTO 
	student_course VALUES (?, ?)
	`

	_, err := utils.DB.Query(
		query,
		studentCourseSelect.StudentID,
		studentCourseSelect.CourseID,
	)

	return err
}

func GetStudentIDByUserID(studentID int) (int, error) {
	var id int
	query := `
	SELECT id from student WHERE user_id = (?)
	`
	err := utils.DB.QueryRow(query, studentID).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return id, nil
}
