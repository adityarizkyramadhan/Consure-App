package transactionimpl

import (
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
