package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"

	"gorm.io/gorm"
)

type TokoRepository interface {
	FindByEmail(ctx context.Context, email string) (daos.User, error)
	FindByUserID(ctx context.Context, id uint) (daos.User, error)
}

type tokoRepository struct {
	db *gorm.DB
}
