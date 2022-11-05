package utils

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"kma_score_api/middlewares"
	"time"
)

func SetUpCron() {
	_, _, err := middlewares.Logger()
	//Implement Cron

	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")

	s := gocron.NewScheduler(location)

	_, err = s.Every(5).Second().Do(func() {
		fmt.Println("INTERVAL 5S")
	})

	if err != nil {
		fmt.Print("ERROR")
		return
	}

	s.StartAsync()
}
