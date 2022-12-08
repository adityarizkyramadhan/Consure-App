package dto

import (
	"Consure-App/domain"

	"gorm.io/gorm"
)

type Expert struct {
	gorm.Model
	Email      string          `json:"email"`
	Nama       string          `json:"nama"`
	Tag        string          `json:"tag"`
	Price      int             `json:"price"`
	Experience string          `json:"experience"`
	Education  string          `json:"education"`
	LinkImage  string          `json:"link_image"`
	Reviews    []domain.Review `json:"reviews"`
}

type DataExpertWithReview struct {
	Expert      *domain.Expert   `json:"expert"`
	AverageStar float64          `json:"average_star"`
	Reviews     []*domain.Review `json:"reviews"`
}

type History struct {
	Expert       *domain.Expert      `json:"expert"`
	Transactions *domain.Transaction `json:"transactions"`
}
