package models

import (
	"fmt"
	"os"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB

func SetupDatabaseConnection() {
	dsn := os.Getenv("GOOSE_DBSTRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database | Error %v", err))
	}

	DB = database
}
