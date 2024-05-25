// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
)

type DaftarAkun struct {
	ID            sql.NullInt64 `json:"id"`
	IDPelanggan   int64         `json:"id_pelanggan"`
	NomorRekening string        `json:"nomor_rekening"`
	Saldo         int64         `json:"saldo"`
}

type DaftarPelanggan struct {
	ID      int64  `json:"id"`
	Nama    string `json:"nama"`
	Nik     string `json:"nik"`
	NomorHp string `json:"nomor_hp"`
	Pin     string `json:"pin"`
}

type DaftarTransaksi struct {
	ID             int64  `json:"id"`
	NomorRekening  string `json:"nomor_rekening"`
	JenisTransaksi string `json:"jenis_transaksi"`
	Nominal        int64  `json:"nominal"`
}

type Mutasi struct {
	ID             int64  `json:"id"`
	NomorRekening  string `json:"nomor_rekening"`
	JenisTransaksi string `json:"jenis_transaksi"`
	Nominal        int64  `json:"nominal"`
}
