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
		return
	}
	database.Connect()
	app := fiber.New(fiber.Config{})

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.ApiResponse(200, "KMA Score API is working very hard", fiber.Map{}))
	})

	app.Get("scores/:id", handlers.GetScoresByStudentCode)

	app.Get("avg-score/:id", handlers.CalculateAvgScore)

	app.Get("/subjects", handlers.AllSubject)

	app.Post("/add-score/:id", handlers.AddScore)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(utils.ApiResponse(404, "Not found", fiber.Map{}))
	})

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
