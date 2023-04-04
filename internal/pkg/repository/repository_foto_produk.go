package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"

	"gorm.io/gorm"
)

type FotoProdukRepositoryImpl struct {
	db *gorm.DB
}

func NewFotoProdukRepository(db *gorm.DB) *FotoProdukRepositoryImpl {
	return &FotoProdukRepositoryImpl{db}
}

func (r *FotoProdukRepositoryImpl) BulkInsert(ctx context.Context, fotoProduks []daos.FotoProduk) ([]daos.FotoProduk, error) {
	err := r.db.Debug().WithContext(ctx).CreateInBatches(&fotoProduks, len(fotoProduks)).Error
	if err != nil {
		return fotoProduks, err
	}

	return fotoProduks, nil
}
