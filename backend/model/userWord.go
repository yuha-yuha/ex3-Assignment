package model

import "gorm.io/gorm"

type UserWord struct {
	gorm.Model
	Japanese string `gorm:"not null"`
	English  string `gorm:"not null"`
	UserID   uint
}
