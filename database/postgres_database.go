package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgres() (*gorm.DB, error) {

	dsn := os.Getenv("POSTGRES_DSN")

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}