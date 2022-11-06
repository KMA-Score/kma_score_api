package utils

import (
	"github.com/go-co-op/gocron"
	"kma_score_api/database"
	"kma_score_api/middlewares"
	"kma_score_api/models"
	"log"
	"time"
)

func _doCron() {
	log.Printf("Executing cron update meilisearch at %v", time.Now())

	// Copy from arahiko-ayami
	var students []models.Student
	database.DBConn.Model(&models.Student{}).Find(&students)

	_, _, err := middlewares.Logger()

	_, err = MeilisearchClient.Index("students").UpdateDocuments(students)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update meilisearch done!")
}

func InitCron() {
	_, _, err := middlewares.Logger()

	// Should be UTC but I like UTC+7
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")

	s := gocron.NewScheduler(location)

	// Main function
	// Cronjob every 15 minutes
	_, err = s.Every(15).Minute().Do(func() {
		_doCron()
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// uhh start async?
	s.StartAsync()
}
