package expert

import (
	"mime/multipart"
)

type ExpertUsecase interface {
	SignUp(input *InputExpert) error
}

type InputExpert struct {
	Email      string                `form:"email" binding:"required"`
	Nama       string                `form:"nama" binding:"required"`
	Price      int                   `form:"price" binding:"required"`
	Experience string                `form:"experience" binding:"required"`
	Education  string                `form:"education" binding:"required"`
	Avatar     *multipart.FileHeader `form:"avatar" binding:"required"`
}
