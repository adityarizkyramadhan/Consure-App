package reviewimpl

import (
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
