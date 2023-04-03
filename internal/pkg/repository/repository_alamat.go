package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"

	"gorm.io/gorm"
)

type AlamatRepository interface {
	FindAllAlamat(ctx context.Context, user daos.User) ([]*daos.Alamat, error)
	FindAlamatByID(ctx context.Context, id uint, user daos.User) (daos.Alamat, error)
	CreateAlamat(ctx context.Context, alamat daos.Alamat) (uint, error)
	UpdateAlamat(ctx context.Context, alamat daos.Alamat) error
	DeleteAlamat(ctx context.Context, id uint) error
}

type alamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &alamatRepository{db}
}

func (r *alamatRepository) FindAllAlamat(ctx context.Context, user daos.User) ([]*daos.Alamat, error) {
	var alamats []*daos.Alamat

	err := r.db.WithContext(ctx).
		Preload("User").
		Where("user_id = ?", user.ID).
		Find(&alamats).Error

	if err != nil {
		return nil, err
	}
	return alamats, nil
}

func (r *alamatRepository) FindAlamatByID(ctx context.Context, id uint, user daos.User) (daos.Alamat, error) {
	var alamat daos.Alamat
	err := r.db.Debug().WithContext(ctx).Preload("User").Where("user_id = ?", user.ID).First(&alamat, id).Error
	if err != nil {
		return daos.Alamat{}, err
	}
	return alamat, nil
}

func (r *alamatRepository) CreateAlamat(ctx context.Context, alamat daos.Alamat) (uint, error) {
	err := r.db.Debug().WithContext(ctx).Create(&alamat).Error

	if err != nil {
		return 0, err
	}

	return alamat.ID, nil
}

func (r *alamatRepository) UpdateAlamat(ctx context.Context, alamat daos.Alamat) error {
	err := r.db.Debug().WithContext(ctx).Model(&alamat).Where("id = ?", alamat.ID).Updates(&alamat).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *alamatRepository) DeleteAlamat(ctx context.Context, id uint) error {
	err := r.db.Debug().WithContext(ctx).Delete(&daos.Alamat{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
