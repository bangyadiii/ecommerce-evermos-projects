package daos

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	UserID           uint   `json:"id_user" gorm:"id_user"`
	AlamatPengiriman uint   `json:"alamat_pengiriman"`
	Alamat           Alamat `gorm:"constraint:OnDelete:NULL"`
	HargaTotal       uint   `json:"harga_total"`
	KodeInvoice      string `json:"kode_invoice"`
	MetodeBayar      string `json:"metode_bayar"`
}

type DetailTrx struct {
	gorm.Model
	TrxID       uint `json:"id_trx" gorm:"id_gorm"`
	Trx         Trx
	LogProdukID uint `json:"id_log_produk"`
	LogProduk   LogProduk
	Kuantitas   int  `json:"kuantitas"`
	HargaTotal  uint `json:"harga_total"`
}
