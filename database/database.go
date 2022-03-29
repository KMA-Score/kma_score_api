package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("kma_score.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DBConn = db
}
