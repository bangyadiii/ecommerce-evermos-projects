package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"fmt"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	GetAllProduks(ctx context.Context, params daos.FilterProduk) (res []*daos.Produk, err error)
	GetProdukByID(ctx context.Context, id uint) (res daos.Produk, err error)
	CreateProduk(ctx context.Context, data daos.Produk) (res uint, err error)
	UpdateProdukByID(ctx context.Context, id uint, data daos.Produk) (res string, err error)
	DeleteProdukByID(ctx context.Context, id uint) (res string, err error)
}

type produkRepositoryImpl struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkRepositoryImpl{db}
}

func (alr *produkRepositoryImpl) GetAllProduks(ctx context.Context, params daos.FilterProduk) (res []*daos.Produk, err error) {
	db := alr.db

	filter := map[string][]any{
		"nama_produk like ? AND deskripsi like ? AND harga_konsumen = ?": []any{fmt.Sprint("%" + params.Nama + "%"), fmt.Sprint("%" + params.Deskripsi + "%"), params.HargaKonsumen},
	}

	// if params.Title != "" {
	// 	db = db.Where("title like ?", "%"+params.Title)
	// }

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	// db = db.Where(map[string]interface{}{"created_at BETWEEN ? AND ?": []string{"2000-01-01 00:00:00", "2000-01-01 00:00:00"}})

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *produkRepositoryImpl) GetProdukByID(ctx context.Context, id uint) (res daos.Produk, err error) {
	if err := alr.db.WithContext(ctx).First(&res, id).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *produkRepositoryImpl) CreateProduk(ctx context.Context, data daos.Produk) (res uint, err error) {
	result := alr.db.WithContext(ctx).Create(&data)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (alr *produkRepositoryImpl) UpdateProdukByID(ctx context.Context, id uint, data daos.Produk) (res string, err error) {
	var produk daos.Produk
	produk, err = alr.GetProdukByID(ctx, id)
	if err != nil {
		return "Update produk failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(produk).Where("id = ? ", id).Updates(&data).Error; err != nil {
		return "Update produk failed", err
	}

	return res, nil
}

func (alr *produkRepositoryImpl) DeleteProdukByID(ctx context.Context, id uint) (res string, err error) {
	var dataProduk daos.Produk
	dataProduk, err = alr.GetProdukByID(ctx, id)

	if err != nil {
		return "Update produk failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataProduk).Delete(&dataProduk).Error; err != nil {
		return "Delete produk failed", err
	}

	return res, nil
}
