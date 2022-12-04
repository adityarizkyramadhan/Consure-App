package transactionimpl

import (
	"Consure-App/domain"
	"Consure-App/repository/transaction"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transaction.TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}

func (repo *TransactionRepositoryImpl) History(id int, data *[]*domain.Transaction) error {
	return repo.DB.Where("id_user = ?", id).Find(data).Error
}
