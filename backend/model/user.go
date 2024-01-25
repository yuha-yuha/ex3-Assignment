package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string     `gorm:"not null" json:"user_name"`
	Email     string     `gorm:"unique" json:"-"`
	Password  string     `gorm:"not null" json:"-"`
	UserWords []UserWord `json:"-"`
}
