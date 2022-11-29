package domain

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	IdUser   int
	IdExpert int
	Review   string
	Star     int
	Komentar string
}
