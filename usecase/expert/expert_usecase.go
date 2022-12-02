package expert

import (
	"Consure-App/domain"
	"mime/multipart"
)

type ExpertUsecase interface {
	SignUp(*InputExpert) error
	FindAll(*[]*domain.Expert) error
	FindById(int, interface{}) error
}

type InputExpert struct {
	Email      string                `form:"email" binding:"required"`
	Nama       string                `form:"nama" binding:"required"`
	Price      int                   `form:"price" binding:"required"`
	Experience string                `form:"experience" binding:"required"`
	Education  string                `form:"education" binding:"required"`
	Tag        string                `form:"tag" binding:"required"`
	Avatar     *multipart.FileHeader `form:"avatar" binding:"required"`
}
