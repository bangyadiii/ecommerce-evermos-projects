package daos

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	UserID       uint   `json:"id_user" gorm:"id_user"`
	Judul        string `json:"judul" gorm:"judul"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
	User User
}

