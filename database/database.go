package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	var dsn string

	if dsn = os.Getenv("DB_PATH"); dsn == "" {
		log.Fatal("DB_PATH is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DBConn = db
}
