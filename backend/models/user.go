package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `json:"email"`
	Nama          string `json:"nama"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	AlamatLengkap string `json:"alamat_lengkap"`
	Status        string `json:"status"`
}
