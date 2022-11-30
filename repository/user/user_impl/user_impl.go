package userimpl

import (
	"Consure-App/repository/user"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (ur *UserRepositoryImpl) FindByUsername(username string, data interface{}) error {
	return ur.DB.Where("username = ?", username).Take(data).Error
}
