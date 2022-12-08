package userimpl

import (
	"Consure-App/domain"
	"Consure-App/middleware"
	"Consure-App/repository/general"
	repouser "Consure-App/repository/user"
	"Consure-App/sdk/upload"
	usecaseuser "Consure-App/usecase/user"
	"fmt"
	"mime/multipart"

	hash "github.com/adityarizkyramadhan/sdk-golang/hash"
)

type UserUsecaseImpl struct {
	RepoUser    repouser.UserRepository
	RepoGeneral general.GeneralRepository
}

func NewUserUsecaseImpl(repoUser repouser.UserRepository, repoGeneral general.GeneralRepository) usecaseuser.UserUsecase {
	return &UserUsecaseImpl{
		RepoUser:    repoUser,
		RepoGeneral: repoGeneral,
	}
}

func (uc *UserUsecaseImpl) SignIn(username, password string) (string, error) {
	user := new(domain.User)
	if err := uc.RepoUser.FindByEmail(username, &user); err != nil {
		return "", err
	}
	hashing := hash.NewPasswordSha512()
	isLogin, err := hashing.ComparePassword(password, user.Password)
	if err != nil {
		return "", err
	}
	if !isLogin {
		return "", fmt.Errorf("password doesn't match")
	}
	token, err := middleware.GenerateJWToken(int(user.ID))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UserUsecaseImpl) SignUp(username, password, email string) (string, error) {
	hashing := hash.NewPasswordSha512()
	passwordHash, err := hashing.GeneratePasswordSha512(password)
	fmt.Println(passwordHash)
	if err != nil {
		return "", err
	}
	user := &domain.User{
		Username: username,
		Password: passwordHash,
		Email:    email,
	}
	if err := uc.RepoGeneral.Create(user); err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(int(user.ID))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UserUsecaseImpl) UpdateProfile(id int, avatar *multipart.FileHeader) (string, error) {
	link, err := upload.UploadImage(avatar)
	if err != nil {
		return "", err
	}
	err = uc.RepoUser.UpdateProfile(id, link)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (uc *UserUsecaseImpl) GetProfile(id int, data *domain.User) error {
	return uc.RepoGeneral.FindById(id, data)
}
