package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"path/filepath"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	dir, err := filepath.Abs("./data/kma_score.db")

	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(sqlite.Open(dir), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DBConn = db
}
