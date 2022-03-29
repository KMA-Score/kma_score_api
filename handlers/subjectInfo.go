package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
)

func AllSubject(c *fiber.Ctx) error {
	var subjects []models.SubjectInfo
	database.DBConn.Model(&models.SubjectInfo{}).Find(&subjects)

	return c.Status(200).JSON(subjects)
}
