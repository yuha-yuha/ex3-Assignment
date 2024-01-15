package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBconn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ex3"))
	if err != nil {
		panic("field to connect database")
	}

	return db
}
