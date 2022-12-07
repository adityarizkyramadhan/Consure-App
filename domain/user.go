package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Email     string `gorm:"unique"`
	Password  string
	LinkImage string
}
