package reviewimpl

import (
	"Consure-App/domain"
	"Consure-App/repository/review"

	"gorm.io/gorm"
)

type ReviewRepositoryImpl struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) review.ReviewRepository {
	return &ReviewRepositoryImpl{
		DB: db,
	}
}
func (repImpl *ReviewRepositoryImpl) FindByIdExpert(id int, data *[]*domain.Review) error {
	return repImpl.DB.Where("id_expert = ?", id).Find(data).Error
}

func (repImpl *ReviewRepositoryImpl) FindByIdUser(id int, data *[]*domain.Review) error {
	return repImpl.DB.Where("id_user = ?", id).Find(data).Error
}
