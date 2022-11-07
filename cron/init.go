package cron

import (
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func InitCron() {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")

	s := gocron.NewScheduler(location)

	// Add new cron job here
	_, err = s.Every(1).Hour().Do(UpdateMeilisearchDocuments)

	if err != nil {
		log.Fatal(err)
	}

	s.StartAsync()
}
