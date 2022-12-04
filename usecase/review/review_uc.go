package review

import (
	"Consure-App/domain"
	"Consure-App/dto"
)

type ReviewUsecase interface {
	FindAll(data *[]*domain.Review) error
	FindById(id int, data interface{}) error
	FindByIdExpert(id int, data *dto.DataExpertWithReview) error
	FindByIdUser(id int, data *[]*domain.Review) error
	Create(int, *InputReview) error
}

type InputReview struct {
	Star     int    `json:"star" binding:"required"`
	Komentar string `json:"komentar" binding:"required"`
	IdExpert int
}
