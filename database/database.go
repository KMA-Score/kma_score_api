package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	var dir string
	var err error

	if dir = os.Getenv("DB_PATH"); dir == "" {
		dir, err = filepath.Abs("./data/kma_score.db")

		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := gorm.Open(sqlite.Open(dir), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DBConn = db
}
