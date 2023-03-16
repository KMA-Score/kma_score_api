package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
	"strings"
)

func GetAllSubjects(c *fiber.Ctx) error {
	type Result struct {
		ID              string `gorm:"column:Id" json:"id"`
		Name            string `gorm:"column:Name" json:"name"`
		NumberOfCredits int    `gorm:"column:NumberOfCredits" json:"numberOfCredits"`
	}

	var subjects []Result
	database.DBConn.Model(&models.Subject{}).Find(&subjects)

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", subjects))
}

func SubjectStatistics(c *fiber.Ctx) error {
	SubjectId := strings.ToUpper(c.Params("SubjectId"))
	var result []models.Score

	database.DBConn.Model(&models.Score{}).Where(&models.Score{SubjectId: SubjectId}).Find(&result)

	var failedStudents = 0
	var passedStudents = 0

	for _, score := range result {
		if !utils.IsPassedSubject(score) {
			failedStudents = failedStudents + 1
		}

		passedStudents = passedStudents + 1
	}

	var data fiber.Map

	if len(result) == 0 {
		data = fiber.Map{}
		return c.Status(404).JSON(utils.ApiResponse(404, "Subject is not exist", data))
	}

	data = fiber.Map{
		"failedStudents": failedStudents,
		"passedStudents": passedStudents,
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", data))
}
