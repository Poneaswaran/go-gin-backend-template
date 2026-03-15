package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {

	db, err := gorm.Open(sqlite.Open("test.db?_pragma=busy_timeout(5000)"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Database connected 🚀")
}
