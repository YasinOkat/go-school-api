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
	Id          string     `gorm:"primary_key;auto_increment" json:"id"`
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
