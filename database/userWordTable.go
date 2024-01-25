package database

import (
	"ej-ex3-backend/model"
)

func CreateWord(word model.UserWord, user model.User) {
	db := DBconn()

	db.Create(&word)
}

func GetALLUsersWords() ([]model.UserWord, error) {
	db := DBconn()

	var uw []model.UserWord

	err := db.Find(&uw).Error

	return uw, err
}

func GetALLUserWords(uid int64) ([]model.UserWord, error) {
	db := DBconn()
	user := model.User{}
	err := db.First(&user, uid).Error

	return user.UserWords, err
}
