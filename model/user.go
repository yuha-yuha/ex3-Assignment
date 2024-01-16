package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"user_name"`
	Email    string `gorm:"unique" json:"user_email"`
	Password string `gorm:"not null" json:"user_password"`
	UserWord []UserWord
}
