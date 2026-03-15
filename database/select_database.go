package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbType := os.Getenv("DB_TYPE")

	var db *gorm.DB
	var err error

	switch dbType {
	case "postgres":
		db, err = connectPostgres()
	case "mysql":
		db, err = connectMySQL()
	default:
		db, err = connectSQLite()
	}

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Database connected 🚀")
}
