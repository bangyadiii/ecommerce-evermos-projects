package daos

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Nama         string  `json:"nama" gorm:"type:text"`
		Email        string  `json:"email" gorm:"type:varchar(255);unique"`
		NoTelp       string  `json:"no_telp" gorm:"type:varchar(20);unique"`
		TanggalLahir *string `json:"tanggal_Lahir" gorm:"type:text"`
		Tentang      *string `json:"tentang" gorm:"type:text"`
		Pekerjaan    *string `json:"pekerjaan" gorm:"type:text"`
		KataSandi    string  `json:"kata_sandi" gorm:"type:text"`
		ProvinsiID   string  `json:"id_provinsi" gorm:"type:text"`
		KotaID       string  `json:"id_kota" gorm:"type:text"`
		IsAdmin      bool    `json:"is_admin" gorm:"index"`
	}

	FilterUser struct {
		Email  string `json:"email"`
		NoTelp string `json:"no_telp"`
	}
)
