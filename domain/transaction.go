package domain

import (
	"time"
)

// id pembayaran
// paket yg dipilih
// jadwal yg dipilih
// total
// harga paket x jumlah
// harga biaya admin
// metode pembayaran : bank/ewallet

type Transaction struct {
	// gorm.Model
	// DeletedAt          DeletedAt `gorm:"index"`
	ID                 uint `gorm:"primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Paket              string    `json:"paket"`
	Jadwal             string    `json:"jadwal"`
	DeadlinePembayaran time.Time `json:"deadline_pembayaran"`
	TotalBeli          int       `json:"total_beli"`
	TotalHarga         int       `json:"total_harga"`
	HargaAdmin         int       `json:"harga_admin"`
	Status             string    `json:"status"`
	MetodePembayaran   string    `json:"metode_pembayaran"`
	IdUser             int       `json:"id_user"`
	IdExpert           int       `json:"id_expert"`
	HargaPaket         int       `json:"harga_paket"`
}
