package generalimpl

import (
	"Consure-App/repository/general"

	"gorm.io/gorm"
)

type GeneralRepositoryImpl struct {
	DB *gorm.DB
}

func NewGeneralRepositoryImpl(db *gorm.DB) general.GeneralRepository {
	return &GeneralRepositoryImpl{
		DB: db,
	}
}

func (g *GeneralRepositoryImpl) Create(data interface{}) error {
	return g.DB.Create(data).Error
}

func (g *GeneralRepositoryImpl) FindById(id int, data interface{}) error {
	return g.DB.First(data, id).Error
}

func (g *GeneralRepositoryImpl) FindAll(data interface{}) error {
	return g.DB.Find(data).Error
}

func (g *GeneralRepositoryImpl) Delete(id int, data interface{}) error {
	return g.DB.Model(data).Delete(" ID = ?", id).Error
}

func (g *GeneralRepositoryImpl) Update(id int, data interface{}) error {
	return g.DB.Where("id = ?", id).Save(data).Error
}
