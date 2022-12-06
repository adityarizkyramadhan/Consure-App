package domain

import "gorm.io/gorm"

type Expert struct {
	gorm.Model
	Nama       string `json:"nama"`
	Tag        string `json:"tag"`
	Price      int    `json:"price"`
	Experience string `json:"experience"`
	LinkImage  string `json:"link_image"`
}
