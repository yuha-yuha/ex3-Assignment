package database

import (
	"ej-ex3-backend/model"
)

func CreateWord(word model.UserWord) {
	db := DBconn()

	db.Create(&word)
}

func GetALLUserWords() ([]model.UserWord, error) {
	db := DBconn()

	var uw []model.UserWord

	err := db.Find(&uw).Error

	return uw, err
}
