package mysql

import (
	"ecommerce-evermos-projects/internal/daos"
	"ecommerce-evermos-projects/internal/helper"
	"fmt"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&daos.User{},
		&daos.Toko{},
		&daos.Alamat{},
		&daos.Category{},
		&daos.Produk{},
		&daos.FotoProduk{},
		&daos.Trx{},
		&daos.LogProduk{},
		&daos.DetailTrx{},
	)

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
