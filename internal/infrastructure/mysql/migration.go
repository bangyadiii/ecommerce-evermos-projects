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

	// var count int64
	// if mysqlDB.Migrator().HasTable(&daos.Book{}) {
	// 	mysqlDB.Model(&daos.Book{}).Count(&count)
	// 	if count < 1 {
	// 		mysqlDB.CreateInBatches(booksSeed, len(booksSeed))
	// 	}
	// }

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
