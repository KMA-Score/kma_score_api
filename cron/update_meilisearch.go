package cron

import (
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
	"log"
	"time"
)

func UpdateMeilisearchDocuments() {
	log.Printf("Executing cron update meilisearch at %v", time.Now())

	var students []models.Student
	database.DBConn.Model(&models.Student{}).Find(&students)

	_, err := utils.MeilisearchClient.Index("students").UpdateDocuments(students)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated meilisearch documents!")
}
