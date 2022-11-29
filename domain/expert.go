package domain

import "gorm.io/gorm"

type Expert struct {
	gorm.Model
	Email      string
	Password   string
	Nama       string
	Price      int
	Experience string
	Education  string
	LinkImage  string
}
