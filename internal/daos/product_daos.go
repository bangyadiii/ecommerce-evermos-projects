package daos

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	Nama          string       `json:"nama_produk"`
	Slug          string       `json:"slug" gorm:"slug;type:text" `
	HargaReseller int          `json:"harga_reseller" gorm:"type:Int"`
	HargaKonsumen int          `json:"harga_konsumen" gorm:"type:Int"`
	Stok          int          `json:"stok" gorm:"type:int"`
	Deskripsi     string       `json:"deskripsi"`
	TokoID        uint         `json:"id_toko" gorm:"column:id_toko;foreignKey:id_toko"`
	Toko          Toko         ``
	FotoProduks   []FotoProduk `json:"produks"`
	CategoryID    uint         `json:"category_id" gorm:"column:category_id;foreignKey:category_id"`
	Category      Category     ``
}

type FotoProduk struct {
	gorm.Model
	ProdukID uint   `json:"id_produk" gorm:"foreignKey:id_produk"`
	Produk   Produk ``
}

type LogProduk struct {
	gorm.Model
	ProdukID      uint         `json:"id_produk" gorm:"column:id_produk;foreignKey:id_produk"`
	Produk        Produk       ``
	Nama          string       `json:"nama_produk" gorm:"type:varchar(255)"`
	Slug          string       `json:"slug" gorm:"slug;type:varchar(255)"`
	HargaReseller int          `json:"harga_reseller" gorm:"int"`
	HargaKonsumen int          `json:"harga_konsumen" gorm:"int"`
	Stok          int          `json:"stok" gorm:"type:int"`
	Deskripsi     string       `json:"deskripsi"`
	TokoID        uint         `json:"id_toko" gorm:"column:id_toko;foreignKey:id_toko"`
	Toko          Toko         ``
	FotoProduks   []FotoProduk `json:"foto_produks" gorm:"-"`
	CategoryID    uint         `json:"id_category" gorm:"column:id_category;foreignKey:id_category"`
	Category      Category     ``
}
