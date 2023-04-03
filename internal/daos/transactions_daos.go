package daos

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	UserID      uint   `json:"id_user" gorm:"foreignKey:id_user"`
	AlamatID    uint   `json:"alamat_id" gorm:"column:alamat_pengiriman;foreignKey:Alamat;references:ID"`
	Alamat      Alamat ``
	HargaTotal  int    `json:"harga_total" gorm:"type:int"`
	KodeInvoice string `json:"kode_invoice" gorm:"type:varchar(255)"`
	MetodeBayar string `json:"metode_bayar" gorm:"type:varchar(100)"`
}

type DetailTrx struct {
	gorm.Model
	TrxID       uint      `json:"id_trx" gorm:"foreignKey:id_trx"`
	Trx         Trx       ``
	LogProdukID uint      `json:"id_log_produk" gorm:"foreignkey:id_log_produk"`
	LogProduk   LogProduk ``
	Kuantitas   int       `json:"kuantitas" gorm:"type:int"`
	HargaTotal  int      `json:"harga_total" gorm:"type:int"`
}
