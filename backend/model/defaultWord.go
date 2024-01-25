package model

import (
	"gorm.io/gorm"
)

type DefaultWord struct {
	gorm.Model
	Japanese string `gorm:"not null"`
	English  string `gorm:"not null"`
}
