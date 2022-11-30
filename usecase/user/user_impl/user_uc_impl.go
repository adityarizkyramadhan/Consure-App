package userimpl

import (
	"Consure-App/domain"
	"Consure-App/middleware"
	"Consure-App/repository/general"
	repouser "Consure-App/repository/user"
	usecaseuser "Consure-App/usecase/user"
	"fmt"
	"log"

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
	if err := uc.RepoUser.FindByUsername(username, &user); err != nil {
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
	if err != nil {
		return "", err
	}
	user := &domain.User{
		Username: username,
		Password: passwordHash,
		Email:    email,
	}
	if err := uc.RepoGeneral.Create(user); err != nil {
		return "", nil
	}
	token, err := middleware.GenerateJWToken(int(user.ID))
	log.Println(user)
	if err != nil {
		return "", err
	}
	return token, nil
}