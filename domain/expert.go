package domain

import "gorm.io/gorm"

type Expert struct {
	gorm.Model
	Email      string `json:"email"`
	Nama       string `json:"nama"`
	Tag        string `json:"tag"`
	Price      int    `json:"price"`
	Experience string `json:"experience"`
	Education  string `json:"education"`
	LinkImage  string `json:"link_image"`
}
