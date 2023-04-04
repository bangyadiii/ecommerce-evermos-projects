package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"
	"fmt"

	"gorm.io/gorm"
)

type TokoRepository interface {
	FindTokoByID(ctx context.Context, ID uint) (daos.Toko, error)
	FindTokoByUserID(ctx context.Context, UserID uint) (daos.Toko, error)
	GetToko(ctx context.Context, params daos.FilterToko) ([]daos.Toko, error)
	UpdateToko(ctx context.Context, id uint, data daos.Toko) (res string, err error)
}

type tokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db}
}

func (r *tokoRepository) FindTokoByID(ctx context.Context, id uint) (data daos.Toko, err error) {
	err = r.db.Debug().WithContext(ctx).First(&data, id).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *tokoRepository) FindTokoByUserID(ctx context.Context, userID uint) (data daos.Toko, err error) {
	err = r.db.Debug().WithContext(ctx).Where("id_user = ?", userID).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *tokoRepository) GetToko(ctx context.Context, params daos.FilterToko) (res []daos.Toko, err error) {
	filter := map[string][]any{
		"nama_toko like ?": []any{fmt.Sprint("%" + params.Name + "%")},
	}

	db := r.db.Debug().WithContext(ctx)
	for key, val := range filter {
		db = db.Where(key, val...)
	}
	if err := db.Limit(params.Limit).Offset(params.Offset).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *tokoRepository) UpdateToko(ctx context.Context, id uint, data daos.Toko) (res string, err error) {
	err = r.db.Debug().WithContext(ctx).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (alr *BooksRepositoryImpl) DeleteTokoByID(ctx context.Context, id uint) (res string, err error) {
	var toko = daos.Toko{}

	if err := alr.db.Model(&toko).Delete(&toko, id).Error; err != nil {
		return "Delete book failed", err
	}

	return res, nil
}
