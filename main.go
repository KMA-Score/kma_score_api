package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"kma_score_api/database"
	"kma_score_api/handlers"
	"kma_score_api/utils"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	database.Connect()
	app := fiber.New(fiber.Config{})

	// create logs directory
	err = utils.CreateDirIfNotExist("./logs")
	if err != nil {
		log.Fatal(err)
	}

	// middlewares
	app.Use(utils.LogToFile)
	app.Use(utils.LogToTerminal)

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.ApiResponse(200, "KMA Score API is working very hard", fiber.Map{}))
	})

	app.Get("statistics", handlers.GeneralScoresStatistics)

	app.Get("statistics/student/:studentCode", handlers.StudentStatistics)

	app.Get("statistics/subject/:subjectCode", handlers.SubjectStatistics)

	app.Get("scores/:studentCode", handlers.GetScoresByStudentCode)

	app.Get("avg-score/:studentCode", handlers.CalculateAvgScore)

	app.Get("/subjects", handlers.AllSubject)

	app.Post("/add-score/:studentCode", handlers.AddScore)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(utils.ApiResponse(404, "Not found", fiber.Map{}))
	})

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
