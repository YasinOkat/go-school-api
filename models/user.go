package models

type UserTypeID uint

const (
	Admin    UserTypeID = 3
	Lecturer UserTypeID = 2
	Student  UserTypeID = 1
)

type UserCreate struct {
	Username    string     `json:"username" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	FirstName   string     `json:"first_name" binding:"required"`
	LastName    string     `json:"last_name" binding:"required"`
	PhoneNumber string     `json:"phone_number" binding:"required"`
	Email       string     `json:"email" binding:"required"`
	UserTypeID  UserTypeID `json:"user_type_id" binding:"required"`
	Status      bool       `json:"status,omitempty"`
}

type User struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	UserTypeID  UserTypeID `json:"user_type_id"`
	Status      bool       `json:"status"`
}
