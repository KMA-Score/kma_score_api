package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
)

func GetAllStudents(c *fiber.Ctx) error {
	var students []models.Student
	database.DBConn.Model(&models.Student{}).Find(&students)

	if len(students) == 0 {
		return c.Status(404).JSON(utils.ApiResponse(404, "Not Found", fiber.Map{}))
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", students))
}
