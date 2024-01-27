package model

import "gorm.io/gorm"

type UserWord struct {
	gorm.Model
	Japanese string `gorm:"not null" json:"japanese"`
	English  string `gorm:"not null" json:"english"`
	UserID   uint   `json:"user_id"`
}
