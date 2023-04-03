package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// func RunMigration(mysqlDB *gorm.DB) {
// 	err := mysqlDB.AutoMigrate(
// 		&daos.User{},
// 		&daos.Toko{},
// 		&daos.Alamat{},
// 		&daos.Category{},
// 		&daos.Produk{},
// 		&daos.FotoProduk{},
// 		&daos.Trx{},
// 		&daos.LogProduk{},
// 		&daos.DetailTrx{},
// 	)

// 	// var count int64
// 	// if mysqlDB.Migrator().HasTable(&daos.Book{}) {
// 	// 	mysqlDB.Model(&daos.Book{}).Count(&count)
// 	// 	if count < 1 {
// 	// 		mysqlDB.CreateInBatches(booksSeed, len(booksSeed))
// 	// 	}
// 	// }

// 	if err != nil {
// 		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
// 	}

// 	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
// }

// func RunGoMigrate(db *gorm.DB) {
// 	driver, err := mysql.Withk(db.DB(), &mysql.Config{})
// 	if err != nil {
// 		panic("failed to get mysql driver")
// 	}

// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file:///path/to/migrations",
// 		"mysql", driver,
// 	)
// }

func main() {
	databaseURL := "mysql://root:@tcp(127.0.0.1:3306)/rakamin_intern?charset=utf8mb4&parseTime=True&loc=Local"

	migrationPath := "file://db/migrations"
	
	
	m, err := migrate.New(migrationPath, databaseURL)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Migrations applied successfully")
}

