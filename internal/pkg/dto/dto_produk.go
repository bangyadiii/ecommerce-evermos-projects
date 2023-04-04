package dto

import (
	"ecommerce-evermos-projects/internal/daos"
	"mime/multipart"
)

type ProdukResp struct {
	ID            uint        `json:"id"`
	Nama          string      `json:"nama_produk"`
	HargaReseller int         `json:"harga_reseller"`
	HargaKonsumen int         `json:"harga_konsumen"`
	Stok          int         `json:"stok"`
	Deskripsi     string      `json:"deskripsi"`
	TokoID        uint        `json:"id_toko"`
	Toko          TokoResp    ``
	CategoryID    uint        `json:"category_id"`
	Category      CategoryRes ``
}

func ProdukDaosToDto(produk daos.Produk) ProdukResp {
	return ProdukResp{
		ID:            produk.ID,
		Nama:          produk.Nama,
		HargaReseller: produk.HargaKonsumen,
		HargaKonsumen: produk.HargaKonsumen,
		Stok:          produk.Stok,
		Deskripsi:     produk.Deskripsi,
		TokoID:        produk.TokoID,
	}
}

func ProdukSliceDaosToDto(produks []*daos.Produk) (res []*ProdukResp) {
	for _, v := range produks {
		produk := ProdukDaosToDto(*v)
		res = append(res, &produk)
	}
	return res
}

type ProdukReqCreate struct {
	Nama          string                  `json:"nama_toko" form:"nama_toko" validate:"required"`
	Photo         []*multipart.FileHeader `form:"-"`
	HargaReseller int                     `json:"harga_reseller" validate:"required"`
	HargaKonsumen int                     `json:"harga_konsumen" validate:"required"`
	Stok          int                     `json:"stok" validate:"required,min:0"`
	Deskripsi     string                  `json:"deskripsi" validate:"required"`
	TokoID        uint                    `json:"id_toko" validate:"required"`
	Toko          TokoResp                ``
	CategoryID    uint                    `json:"category_id" validate:"required"`
	Category      CategoryRes             ``
}

type ProdukReqUpdate struct {
	Nama          string                  `json:"nama_toko" form:"nama_toko"`
	Photo         []*multipart.FileHeader `form:"-"`
	HargaReseller int                     `json:"harga_reseller"`
	HargaKonsumen int                     `json:"harga_konsumen"`
	Stok          int                     `json:"stok"`
	Deskripsi     string                  `json:"deskripsi"`
	TokoID        uint                    `json:"id_toko"`
	Toko          TokoResp                ``
	CategoryID    uint                    `json:"category_id"`
	Category      CategoryRes             ``
}

type FilterProduk struct {
	Nama  string `json:"nama_toko"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}
