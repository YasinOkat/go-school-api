package models

type StudentRead struct {
	ID      int  `json:"id"`
	UserID  int  `json:"userID"`
	MajorID int  `json:"majorID"`
	Status  bool `json:"status"`
}

type StudentCreate struct {
	UserID  int  `json:"userID" binding:"required"`
	MajorID int  `json:"majorID" binding:"required"`
	Status  bool `json:"status,omitempty"`
}

type StudentUserCreate struct {
	MajorID     int        `json:"majorID" binding:"required"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	UserTypeID  UserTypeID `json:"user_type_id"`
	Status      bool       `json:"status"`
}

type StudentCourseSelect struct {
	StudentID int `json:"studentID" binding:"required"`
	CourseID  int `json:"courseID" binding:"required"`
}

type Grade struct {
	StudentID int     `json:"studentID" binding:"required"`
	CourseID  int     `json:"courseID" binding:"required"`
	Grade     float64 `json:"grade" binding:"required"`
}

type StudentCourse struct {
	StudentID  int    `json:"studentID" binding:"required"`
	CourseID   int    `json:"courseID" `
	Username   string `json:"username" binding:"required"`
	CourseName string `json:"courseName" `
}
