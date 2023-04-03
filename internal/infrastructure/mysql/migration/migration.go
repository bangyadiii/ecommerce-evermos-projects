package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

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
