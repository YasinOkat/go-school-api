package models

import (
	"time"
)

type UserTypeID uint

const (
	Admin    UserTypeID = 3
	Lecturer UserTypeID = 2
	Student  UserTypeID = 1
)

type UserCreate struct {
	Username    string     `gorm:"unique;not null" json:"username"`
	Password    string     `gorm:"not null" json:"password"`
	FirstName   string     `gorm:"not null" json:"first_name"`
	LastName    string     `gorm:"not null" json:"last_name"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	Email       string     `gorm:"not null" json:"email"`
	UserTypeID  UserTypeID `gorm:"not null" json:"user_type_id"`
	Status      bool       `gorm:"not null;default:true" json:"status"`
}

type User struct {
	Username    string     `gorm:"unique;not null" json:"username"`
	Password    string     `gorm:"not null" json:"password"`
	FirstName   string     `gorm:"not null" json:"first_name"`
	LastName    string     `gorm:"not null" json:"last_name"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	Email       string     `gorm:"not null" json:"email"`
	UserTypeID  UserTypeID `gorm:"not null" json:"user_type_id"`
	Status      bool       `gorm:"not null;default:true" json:"status"`
	CreatedAt   time.Time  `gorm:"-" json:"created_at,omitempty"`
	UpdatedAt   time.Time  `gorm:"-" json:"updated_at,omitempty"`
}

func (User) TableName() string {
	return "user"
}
