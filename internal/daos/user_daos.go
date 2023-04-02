package daos

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Nama         string  `json:"nama" gorm:"nama"`
		Email        string  `json:"email" gorm:"email,index:unique_email,unique"`
		NoTelp       string  `json:"no_telp" gorm:"no_telp,index:unique_no_telp,unique"`
		TanggalLahir *string `json:"tanggal_Lahir" gorm:"tanggal_lahir"`
		Tentang      *string `json:"tentang" gorm:"tentang"`
		Pekerjaan    *string `json:"pekerjaan" gorm:"pekerjaan"`
		Password     string  `json:"kata_sandi" gorm:"kata_sandi"`
		ProvinsiID   string  `json:"id_provinsi" gorm:"id_provinsi"`
		KotaID       string  `json:"id_kota" gorm:"id_kota"`
		IsAdmin      bool    `json:"is_admin"`
	}

	FilterUser struct {
		Email  string `json:"email"`
		NoTelp string `json:"no_telp"`
	}
)
