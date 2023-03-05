package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
	"math"
	"strings"
)

func AddScore(c *fiber.Ctx) error {
	id := strings.ToUpper(c.Params("StudentId"))

	type addScorePayload struct {
		SubjectId            string `json:"SubjectId"`
		FirstComponentScore  string `json:"firstComponentScore"`
		SecondComponentScore string `json:"secondComponentScore"`
		ExamScore            string `json:"examScore"`
		AvgScore             string `json:"avgScore"`
		AlphabetScore        string `json:"alphabetScore"`
	}

	var payload addScorePayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(utils.ApiResponse(400, "Bad Request", fiber.Map{}))
	}

	var score models.Score
	database.DBConn.Model(&models.Score{}).
		Where(&models.Score{SubjectId: payload.SubjectId, StudentId: id}).
		Find(&score)

	if score.StudentId != "" && score.SubjectId != "" {
		return c.Status(409).JSON(utils.ApiResponse(409, "Score is already exist", fiber.Map{}))
	}

	if id != "" {
		newScore := models.Score{
			StudentId:            id,
			SubjectId:            payload.SubjectId,
			FirstComponentScore:  payload.FirstComponentScore,
			SecondComponentScore: payload.SecondComponentScore,
			ExamScore:            payload.ExamScore,
			AvgScore:             payload.AvgScore,
			AlphabetScore:        payload.AlphabetScore,
		}

		if utils.IsValidScore(newScore) {
			if newScore.AlphabetScore == "" {
				var err error
				newScore.AvgScore, err = utils.CalcSubjectAvgScore(newScore)
				newScore.AlphabetScore = utils.ConvertDecimalScoreToAlphabetScore(newScore.AvgScore)
				if err != nil {
					return c.Status(400).JSON(utils.ApiResponse(400, "Can't parse score", fiber.Map{}))
				}
			}

			database.DBConn.Model(&models.Score{}).Create(&newScore)
			return c.Status(200).JSON(utils.ApiResponse(200, "Added score!", fiber.Map{}))
		}
	}

	return c.Status(400).JSON(utils.ApiResponse(400, "Your score is not valid", fiber.Map{}))
}

func GeneralScoresStatistics(c *fiber.Ctx) error {
	var numberOfStudents int64
	var numberOfDebtors int64
	var numberOfSubjects int64

	database.DBConn.Model(&models.Score{}).Group("StudentId").Count(&numberOfStudents)
	database.DBConn.Model(&models.Score{}).
		Where(&models.Score{AlphabetScore: ""}).
		Where(&models.Score{AlphabetScore: "F"}).
		Group("id").
		Count(&numberOfDebtors)
	database.DBConn.Model(&models.Subject{}).Count(&numberOfSubjects)

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", fiber.Map{
		"numberOfStudents": numberOfStudents,
		"numberOfDebtors":  numberOfDebtors,
		"numberOfSubjects": numberOfSubjects,
	}))
}

func StudentStatistics(c *fiber.Ctx) error {
	id := strings.ToUpper(c.Params("StudentId"))
	var result models.Student

	database.DBConn.Where("Id = ?", id).Preload("Scores").Preload("Scores.Subject").Find(&result)

	failedSubjects := 0
	passedSubjects := 0

	for _, score := range result.Scores {
		if !utils.IsPassedSubject(score) {
			failedSubjects = failedSubjects + 1
		} else {
			passedSubjects = passedSubjects + 1
		}
	}

	var sumAvgScore = 0.0
	var totalNumberOfCredit = 0

	var SubjectIds []string
	for _, score := range result.Scores {
		SubjectIds = append(SubjectIds, score.SubjectId)
	}

	var subjects []models.Subject
	database.DBConn.Model(&models.Subject{}).Where("(Id) IN ?", SubjectIds).Find(&subjects)

	for _, score := range result.Scores {
		if utils.ShouldCalculateAverageScore(score) {
			index := utils.GetSubjectIndex(subjects, score)
			if index >= 0 {
				totalNumberOfCredit = totalNumberOfCredit + subjects[index].NumberOfCredits
				sumAvgScore = sumAvgScore + utils.AlphabetScoreToTetraScore(score.AlphabetScore)*float64(subjects[index].NumberOfCredits)
			}
		}
	}

	avgScore := sumAvgScore / float64(totalNumberOfCredit)
	avgScore = math.Round(avgScore*100) / 100

	var data fiber.Map

	if result.ID == "" {
		data = fiber.Map{}
		return c.Status(404).JSON(utils.ApiResponse(404, "Student is not exist", data))
	}

	data = fiber.Map{
		"id":             result.ID,
		"name":           result.Name,
		"class":          result.Class,
		"passedSubjects": passedSubjects,
		"failedSubjects": failedSubjects,
		"scores":         result.Scores,
		"avgScore":       avgScore,
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", data))
}
