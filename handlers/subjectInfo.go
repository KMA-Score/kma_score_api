package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
	"strings"
)

func AllSubject(c *fiber.Ctx) error {
	var subjects []models.SubjectInfo
	database.DBConn.Model(&models.SubjectInfo{}).Find(&subjects)

	return c.Status(200).JSON(subjects)
}

func SubjectStatistics(c *fiber.Ctx) error {
	type Result struct {
		StudentCode          string `gorm:"column:studentCode" json:"studentCode"`
		Name                 string `gorm:"column:name" json:"name"`
		NumberOfCredit       int    `gorm:"column:noc" json:"numberOfCredit"`
		FirstComponentScore  string `gorm:"column:TP1" json:"firstComponentScore"`
		SecondComponentScore string `gorm:"column:TP2" json:"secondComponentScore"`
		ExamScore            string `gorm:"column:THI" json:"examScore"`
		AvgScore             string `gorm:"column:TONGKET" json:"avgScore"`
		AlphabetScore        string `gorm:"column:DIEMCHU" json:"alphabetScore"`
	}

	subjectCode := strings.ToUpper(c.Params("subjectCode"))
	var result []Result

	database.DBConn.Raw("SELECT * FROM studentScore inner join subjectInfo on studentScore.subjectCode = subjectInfo.subjectCode WHERE subjectInfo.subjectCode = ?", subjectCode).
		Find(&result)

	var failedStudents = 0
	var passedStudents = 0

	for _, score := range result {
		if !utils.IsPassedSubject(score.AlphabetScore) {
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
