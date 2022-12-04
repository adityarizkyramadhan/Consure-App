package transactionimpl

import (
	"Consure-App/repository/general"
	transRep "Consure-App/repository/transaction"
	transUc "Consure-App/usecase/transaction"
)

type TransactionUsecase struct {
	GeneralRepo general.GeneralRepository
	TransRepo   transRep.TransactionRepository
}

func NewTransactionUsecase(general general.GeneralRepository, transRepo transRep.TransactionRepository) transUc.TransactionUsecase {
	return &TransactionUsecase{
		GeneralRepo: general,
		TransRepo:   transRepo,
	}
}
