package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectMySQL() (*gorm.DB, error) {

	dsn := os.Getenv("MYSQL_DSN")

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}