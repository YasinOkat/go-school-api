package models

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
