package daos

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	Nama          string       `json:"nama_produk"`
	Slug          string       `json:"slug" gorm:"slug"`
	HargaReseller string       `json:"harga_reseller"`
	HargaKonsumen string       `json:"harga_konsumen"`
	Stok          string       `json:"stok" `
	Deskripsi     string       `json:"deskripsi"`
	TokoID        uint         `json:"id_toko" gorm:"id_toko"`
	Toko          Toko         `gorm:"constraint:OnDelete:CASCADE"`
	FotoProduks   []FotoProduk `json:"foto_produks"`
	CategoryID    uint
	Category      Category `gorm:"constraint:OnDelete:CASCADE"`
}

type FotoProduk struct {
	gorm.Model
	ProdukID uint
	Produk   Produk `gorm:"constraint:OnDelete:CASCADE"`
}

type LogProduk struct {
	gorm.Model
	ProdukID      uint `json:"id_produk" gorm:"id_produk"`
	Produk        Produk
	Nama          string       `json:"nama_produk"`
	Slug          string       `json:"slug" gorm:"slug"`
	HargaReseller string       `json:"harga_reseller"`
	HargaKonsumen string       `json:"harga_konsumen"`
	Stok          string       `json:"stok" `
	Deskripsi     string       `json:"deskripsi"`
	TokoID        uint         `json:"id_toko" gorm:"id_toko"`
	Toko          Toko         `gorm:"constraint:OnDelete:CASCADE"`
	FotoProduks   []FotoProduk `json:"foto_produks"`
	CategoryID    uint
	Category      Category `gorm:"constraint:OnDelete:CASCADE"`
}
