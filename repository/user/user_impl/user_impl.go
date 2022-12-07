package userimpl

import (
	"Consure-App/domain"
	"Consure-App/repository/user"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (ur *UserRepositoryImpl) FindByEmail(email string, data interface{}) error {
	return ur.DB.Where("email = ?", email).Take(data).Error
}

func (ur *UserRepositoryImpl) UpdateProfile(id int, link string) error {
	return ur.DB.Model(&domain.User{}).Where("id = ?", id).Update("link_image", link).Error
}
