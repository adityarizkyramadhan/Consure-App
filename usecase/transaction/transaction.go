package transaction

import "Consure-App/domain"

type TransactionUsecase interface {
	Create(*InputTransaction, int) error
	History(int, *[]*domain.Transaction) error
}

type InputTransaction struct {
	Paket      string `json:"paket"`
	Jadwal     string `json:"jadwal"`
	HargaPaket int    `json:"harga_paket"`
	TotalBeli  int    `json:"total_beli"`
	IdExpert   int    `json:"id_expert"`
}
