package transaction

import (
	"Consure-App/dto"
)

type TransactionUsecase interface {
	Create(*InputTransaction, int) error
	History(int, string, *[]*dto.History) error
}

type InputTransaction struct {
	Paket      string `json:"paket"`
	Jadwal     string `json:"jadwal"`
	HargaPaket int    `json:"harga_paket"`
	TotalBeli  int    `json:"total_beli"`
	IdExpert   int    `json:"id_expert"`
}
