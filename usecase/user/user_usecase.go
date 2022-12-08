package user

import (
	"Consure-App/domain"
	"mime/multipart"
)

type UserUsecase interface {
	SignIn(string, string) (string, error)
	SignUp(string, string, string) (string, error)
	UpdateProfile(int, *multipart.FileHeader) (string, error)
	GetProfile(int, *domain.User) error
}
