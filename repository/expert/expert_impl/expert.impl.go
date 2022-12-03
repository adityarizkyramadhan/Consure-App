package expertimpl

import (
	"Consure-App/domain"
	"Consure-App/repository/expert"
	"fmt"

	"gorm.io/gorm"
)

type ExpertRepoImpl struct {
	DB *gorm.DB
}

func NewExpertRepository(db *gorm.DB) expert.ExpertRepository {
	return &ExpertRepoImpl{DB: db}
}

func (er *ExpertRepoImpl) FindByTag(tag string, data *[]*domain.Expert) error {
	tag = fmt.Sprintln("%", tag, "%")
	return er.DB.Where("tag LIKE ?", tag).Find(data).Error
}
