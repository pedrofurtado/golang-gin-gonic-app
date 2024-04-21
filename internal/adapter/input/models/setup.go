package models

import (
	"os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB

func SetupDBConnection() {
	dsn := os.Getenv("GOOSE_DBSTRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database
}
