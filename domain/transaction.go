package domain

import (
	"time"

	"gorm.io/gorm"
)

// id pembayaran
// paket yg dipilih
// jadwal yg dipilih
// total
// harga paket x jumlah
// harga biaya admin
// metode pembayaran : bank/ewallet

type Transaction struct {
	gorm.Model
	Paket            string    `json:"paket"`
	Jadwal           time.Time `json:"jadwal"`
	TotalBeli        int       `json:"total_beli"`
	TotalHarga       int       `json:"total_harga"`
	HargaAdmin       int       `json:"harga_admin"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	IdUser           int       `json:"id_user"`
	IdExpert         int       `json:"id_expert"`
}
