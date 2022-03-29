package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
	"math"
)

func GetScoresByStudentCode(c *fiber.Ctx) error {
	studentCode := c.Params("id")

	if studentCode != "" {
		var scores []models.StudentScore
		database.DBConn.Model(&models.StudentScore{}).Where(&models.StudentScore{StudentCode: studentCode}).Find(&scores)

		return c.Status(200).JSON(utils.ApiResponse(200, "OK", scores))
	}

	return c.Status(400).JSON(utils.ApiResponse(400, "Bad Request", fiber.Map{}))

}

func CalculateAvgScore(c *fiber.Ctx) error {
	studentCode := c.Params("id")

	if studentCode != "" {
		var scores []models.StudentScore

		database.DBConn.Model(&models.StudentScore{}).
			Where(&models.StudentScore{StudentCode: studentCode}).
			Where("subjectCode NOT LIKE ?", "ATQGTC%").
			Find(&scores)

		var sumAvgScore = 0.0
		var totalNumberOfCredit = 0

		var subjectCodes []string
		for _, element := range scores {
			subjectCodes = append(subjectCodes, element.SubjectCode)
		}

		var subjects []models.SubjectInfo
		database.DBConn.Model(&models.SubjectInfo{}).Where("(subjectCode) IN ?", subjectCodes).Find(&subjects)

		for _, element := range scores {
			index := utils.GetSubjectIndex(subjects, element)
			if index >= 0 {
				totalNumberOfCredit = totalNumberOfCredit + subjects[index].NumberOfCredit
				sumAvgScore = sumAvgScore + utils.AlphabetScoreToTetraScore(element.AlphabetScore)*float64(subjects[index].NumberOfCredit)
			}
		}

		avgScore := sumAvgScore / float64(totalNumberOfCredit)
		avgScore = math.Floor(avgScore*100) / 100

		return c.Status(200).JSON(utils.ApiResponse(200, "OK", fiber.Map{
			"studentCode":         studentCode,
			"avgScore":            avgScore,
			"totalNumberOfCredit": totalNumberOfCredit,
		}))
	}

	return c.Status(400).JSON(utils.ApiResponse(400, "Bad Request", fiber.Map{}))
}

func AddScore(c *fiber.Ctx) error {
	studentCode := c.Params("id")

	type addScorePayload struct {
		SubjectCode   string `json:"subjectCode"`
		AlphabetScore string `json:"alphabetScore"`
	}

	var payload addScorePayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(utils.ApiResponse(400, "Bad Request", fiber.Map{}))
	}

	var score models.StudentScore
	database.DBConn.Model(&models.StudentScore{}).
		Where(&models.StudentScore{SubjectCode: payload.SubjectCode, StudentCode: studentCode}).
		Find(&score)

	if score.StudentCode != "" && score.SubjectCode != "" {
		return c.Status(409).JSON(utils.ApiResponse(409, "Score is already exist", fiber.Map{}))
	}

	if studentCode != "" {
		newScore := models.StudentScore{
			StudentCode:   studentCode,
			SubjectCode:   payload.SubjectCode,
			AlphabetScore: payload.AlphabetScore,
		}
		database.DBConn.Model(&models.StudentScore{}).Create(&newScore)
		return c.Status(200).JSON(utils.ApiResponse(200, "Added score!", fiber.Map{}))
	}

	return c.Status(400).JSON(utils.ApiResponse(400, "Bad Request", fiber.Map{}))
}
