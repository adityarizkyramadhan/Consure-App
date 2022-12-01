package user

import "mime/multipart"

type UserUsecase interface {
	SignIn(string, string) (string, error)
	SignUp(string, string, string) (string, error)
	UpdateProfile(int, *multipart.FileHeader) (string, error)
}
