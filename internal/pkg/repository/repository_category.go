package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAllCategory(ctx context.Context) ([]daos.Category, error)
	CreateCategory(ctx context.Context, data daos.Category) (uint, error)
	FindCategoryByID(ctx context.Context, id uint) (daos.Category, error)
	UpdateCategoryByID(ctx context.Context, id uint, data daos.Category) (string, error)
	DeleteCategoryByID(ctx context.Context, id uint) (res string, err error)
}

type categoryRepoImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepoImpl{db: db}
}

func (r *categoryRepoImpl) FindAllCategory(ctx context.Context) ([]daos.Category, error) {
	var res []daos.Category
	err := r.db.Debug().WithContext(ctx).Find(&res).Error

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *categoryRepoImpl) CreateCategory(ctx context.Context, data daos.Category) (uint, error) {
	err := r.db.Debug().WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *categoryRepoImpl) FindCategoryByID(ctx context.Context, id uint) (daos.Category, error) {
	var res daos.Category
	err := r.db.Debug().WithContext(ctx).First(&res, id).Error

	if err != nil {
		return res, gorm.ErrRecordNotFound
	}

	return res, nil
}

func (r *categoryRepoImpl) UpdateCategoryByID(ctx context.Context, id uint, data daos.Category) (string, error) {
	var res daos.Category

	err := r.db.Debug().WithContext(ctx).Where("id = ?", id).Find(&res).Error
	if err != nil {
		return "Update category failed", gorm.ErrRecordNotFound
	}

	if err := r.db.Where("id = ? ", id).Updates(&data).Error; err != nil {
		return "Update category failed", err
	}

	return "Update category success", nil
}

func (r *categoryRepoImpl) DeleteCategoryByID(ctx context.Context, id uint) (res string, err error) {
	var data daos.Category

	if err := r.db.Debug().WithContext(ctx).Where("id = ?", id).Delete(&data).Error; err != nil {
		return "Delete category failed", err
	}

	return res, nil
}
