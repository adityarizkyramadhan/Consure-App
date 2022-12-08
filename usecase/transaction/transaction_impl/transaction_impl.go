package transactionimpl

import (
	"Consure-App/domain"
	"Consure-App/dto"
	"Consure-App/repository/general"
	transRep "Consure-App/repository/transaction"
	transUc "Consure-App/usecase/transaction"
	"time"
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

func (uc *TransactionUsecase) Create(input *transUc.InputTransaction, idUser int) error {
	data := &domain.Transaction{
		Paket:              input.Paket,
		Jadwal:             input.Jadwal,
		HargaAdmin:         1000,
		HargaPaket:         input.HargaPaket,
		TotalHarga:         (input.HargaPaket * input.TotalBeli) + 1000,
		DeadlinePembayaran: time.Now().Add(15 * time.Minute),
		Status:             "scheduled",
		MetodePembayaran:   "Transfer Bank",
		TotalBeli:          input.TotalBeli,
		IdUser:             idUser,
		IdExpert:           input.IdExpert,
	}
	return uc.GeneralRepo.Create(data)
}

func (uc *TransactionUsecase) History(id int, status string, data *[]*dto.History) error {
	reservasi := []*domain.Transaction{}
	if err := uc.TransRepo.History(id, status, &reservasi); err != nil {
		return err
	}

	if len(reservasi) == 0 {
		return nil
	}

	for i := range reservasi {
		expert := new(domain.Expert)
		if err := uc.GeneralRepo.FindById(reservasi[i].IdExpert, expert); err != nil {
			return err
		}
		his := &dto.History{
			Expert:       expert,
			Transactions: reservasi[i],
		}
		*data = append(*data, his)
	}
	return nil
}
