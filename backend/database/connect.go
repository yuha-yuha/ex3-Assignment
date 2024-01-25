package database

import (
	"ej-ex3-backend/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBconn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ex3"))
	if err != nil {
		panic("field to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserWord{})
	db.AutoMigrate(&model.DefaultWord{})

	return db
}
