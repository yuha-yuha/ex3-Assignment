package model

import (
	"ej-ex3-backend/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	UserWord []UserWord
}

func GetUser(email string) User {
	db := database.DBconn()

	var user User

	db.Where("email = ?", email).First(&user)

	return user
}
