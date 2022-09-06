package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"kma_score_api/database"
	"kma_score_api/handlers"
	"kma_score_api/utils"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	err = utils.CreateDirIfNotExist("./logs")
	LogToFile, LogToTerminal, err := utils.Logger()
	Cors := utils.Cors()

	if err != nil {
		log.Fatal(err)
	}

	database.Connect()
	app := fiber.New(fiber.Config{})

	// middlewares
	app.Use(LogToFile)
	app.Use(LogToTerminal)
	app.Use(Cors)

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(utils.ApiResponse(200, "KMA Score API is working very hard", fiber.Map{}))
	})

	app.Get("statistics", handlers.GeneralScoresStatistics)

	app.Get("student/:StudentId", handlers.StudentStatistics)

	app.Get("subject/:SubjectId", handlers.SubjectStatistics)

	app.Get("/subjects", handlers.AllSubject)

	app.Post("/add-score/:StudentId", handlers.AddScore)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(utils.ApiResponse(404, "Not found", fiber.Map{}))
	})

	var enableSsl, _ = strconv.ParseBool(os.Getenv("ENABLE_SSL"))

	if enableSsl {
		err = app.ListenTLS(":8080", "./cert/public.pem", "./cert/private.pem")
	} else {
		err = app.Listen(":8080")
	}

	if err != nil {
		log.Fatal(err)
	}
}
