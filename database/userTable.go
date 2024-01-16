package database

import (
	"ej-ex3-backend/model"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(email string) model.User {
	db := DBconn()

	var user model.User

	db.Where("email = ?", email).First(&user)

	return user
}

func CreateUser(user model.User) {
	db := DBconn()

	BytesPassword := []byte(user.Password)
	HashedPassword, _ := bcrypt.GenerateFromPassword(BytesPassword, 5)
	user.Password = string(HashedPassword)

	db.Create(&user)

}
