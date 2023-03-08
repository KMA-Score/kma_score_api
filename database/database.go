package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kma_score_api/models"
	"log"
	"os"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")

	if dsn == "" {
		log.Fatal("Database env variables were not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Auto migrate models if not exists
	err = db.AutoMigrate(&models.ApiKey{})

	if err != nil {
		log.Fatal(err)
	}

	DBConn = db
}
