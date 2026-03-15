package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

func connectSQLite() (*gorm.DB, error) {
	dsn := os.Getenv("SQLITE_DSN")
	if dsn == "" {
		dsn = "test.db?_pragma=busy_timeout(5000)"
	}

	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
