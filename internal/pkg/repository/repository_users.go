package repository

import (
	"context"
	"ecommerce-evermos-projects/internal/daos"

	"gorm.io/gorm"
)

type UsersRepository interface {
	SaveUser(ctx context.Context, user daos.User, toko daos.Toko) (daos.User, error)
	FindByEmail(ctx context.Context, email string) (daos.User, error)
	FindByNoTelp(ctx context.Context, noTelp string) (daos.User, error)
	FindByUserID(ctx context.Context, id uint) (daos.User, error)
	UpdateUser(ctx context.Context, user daos.User) (daos.User, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) SaveUser(ctx context.Context, user daos.User, toko daos.Toko) (daos.User, error) {
	db := r.db.Debug().WithContext(ctx).Begin()
	err := db.Create(&user).Error
	if err != nil {
		db.Rollback()
		return user, err
	}
	toko.UserID = user.ID

	err = db.Create(&toko).Error

	if err != nil {
		db.Rollback()
		return user, err
	}

	err = db.Commit().Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersRepository) FindByNoTelp(ctx context.Context, noTelp string) (daos.User, error) {
	var user daos.User
	err := r.db.Debug().WithContext(ctx).Where("no_telp = ?", noTelp).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *usersRepository) FindByEmail(ctx context.Context, email string) (daos.User, error) {
	var user daos.User
	err := r.db.Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *usersRepository) FindByUserID(ctx context.Context, id uint) (daos.User, error) {
	var user daos.User
	err := r.db.Debug().WithContext(ctx).Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *usersRepository) UpdateUser(ctx context.Context, user daos.User) (daos.User, error) {
	err := r.db.Debug().WithContext(ctx).Model(&user).Where("id = ?", user.ID).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
