package daos

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	UserID       uint   `json:"id_user" gorm:"column:id_user;foreignKey:id_usr"`
	Judul        string `json:"judul" gorm:"column:judul;type:varchar(255)"`
	NamaPenerima string `json:"nama_penerima" gorm:"column;type:varchar(255)"`
	NoTelp       string `json:"no_telp" gorm:"column:no_telp;type:varchar(20)"`
	DetailAlamat string `json:"detail_alamat" gorm:"column:detail_alamat;type:varchar(255)"`
	User         User   ``
}
