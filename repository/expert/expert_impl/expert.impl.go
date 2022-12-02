package expertimpl

import (
	"Consure-App/repository/expert"

	"gorm.io/gorm"
)

type ExpertRepoImpl struct {
	DB *gorm.DB
}

func NewExpertRepository(db *gorm.DB) expert.ExpertRepository {
	return &ExpertRepoImpl{DB: db}
}
