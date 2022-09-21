package utils

import (
	"github.com/meilisearch/meilisearch-go"
	"kma_score_api/config"
	"kma_score_api/database"
	"kma_score_api/models"
	"log"
	"os"
)

var MeilisearchClient *meilisearch.Client

func MeilisearchInit() {

	MeilisearchClient = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   "http://" + os.Getenv("MEILISEARCH_HOST") + ":" + os.Getenv("MEILISEARCH_PORT"),
		APIKey: os.Getenv("MEILISEARCH_API_KEY"),
	})

	var students []models.Student
	database.DBConn.Model(&models.Student{}).Find(&students)

	_, err := MeilisearchClient.Index("students").AddDocuments(students)
	_, err = MeilisearchClient.Index("students").UpdateRankingRules(&config.StudentRankingRules)
	_, err = MeilisearchClient.Index("students").UpdateDisplayedAttributes(&config.StudentDisplayedAttributes)

	if err != nil {
		log.Fatal(err)
	}
}
