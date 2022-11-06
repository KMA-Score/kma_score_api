package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"kma_score_api/database"
	"kma_score_api/handlers"
	"kma_score_api/middlewares"
	"kma_score_api/utils"
	"log"
)

func main() {
	err := godotenv.Load()
	err = utils.CreateDirIfNotExist("./logs")
	LogToFile, LogToTerminal, err := middlewares.Logger()
	Cors := middlewares.Cors()
	Limiter := middlewares.Limiter()

	if err != nil {
		log.Fatal(err)
	}

	database.Connect()
	utils.MeilisearchInit()

	// IMPORTANT: cron must be init before http startup and after database + meiliSearch init
	utils.InitCron()

	app := fiber.New(fiber.Config{})

	// middlewares
	app.Use(LogToFile)
	app.Use(LogToTerminal)
	app.Use(Cors)
	app.Use(Limiter)

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.ApiResponse(200, "KMA Score API is working very hard", fiber.Map{}))
	})

	app.Get("statistics", handlers.GeneralScoresStatistics)

	app.Get("student/:StudentId", handlers.StudentStatistics)

	app.Get("subject/:SubjectId", handlers.SubjectStatistics)

	app.Get("/subjects", handlers.AllSubject)

	app.Post("/add-score/:StudentId", handlers.AddScore)

	app.Get("/search/*", handlers.Search)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(utils.ApiResponse(404, "Not found", fiber.Map{}))
	})

	err = app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
