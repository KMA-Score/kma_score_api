package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kma_score_api/database"
	"kma_score_api/models"
	"kma_score_api/utils"
	"math"
	"strings"
)

func GetScoresByStudentCode(c *fiber.Ctx) error {
	studentCode := strings.ToUpper(c.Params("studentCode"))

	if studentCode != "" {
		var scores []models.StudentScore
		database.DBConn.Model(&models.StudentScore{}).Where(&models.StudentScore{StudentCode: studentCode}).Find(&scores)

		return c.Status(200).JSON(utils.ApiResponse(200, "OK", scores))
	}

	return c.Status(400).JSON(utils.ApiResponse(400, "Bad Request", fiber.Map{}))
}

func CalculateAvgScore(c *fiber.Ctx) error {
	studentCode := strings.ToUpper(c.Params("studentCode"))

	if studentCode != "" {
		var scores []models.StudentScore

		database.DBConn.Model(&models.StudentScore{}).
			Where(&models.StudentScore{StudentCode: studentCode}).
			Where("subjectCode NOT LIKE ?", "ATQGTC%").
			Where("NOT (DIEMCHU = \"\")").
			Find(&scores)

		var sumAvgScore = 0.0
		var totalNumberOfCredit = 0

		var subjectCodes []string
		for _, score := range scores {
			subjectCodes = append(subjectCodes, score.SubjectCode)
		}

		var subjects []models.SubjectInfo
		database.DBConn.Model(&models.SubjectInfo{}).Where("(subjectCode) IN ?", subjectCodes).Find(&subjects)

		for _, score := range scores {
			index := utils.GetSubjectIndex(subjects, score)
			if index >= 0 {
				totalNumberOfCredit = totalNumberOfCredit + subjects[index].NumberOfCredit
				sumAvgScore = sumAvgScore + utils.AlphabetScoreToTetraScore(score.AlphabetScore)*float64(subjects[index].NumberOfCredit)
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
	studentCode := strings.ToUpper(c.Params("studentCode"))

	type addScorePayload struct {
		SubjectCode          string `json:"subjectCode"`
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

	var score models.StudentScore
	database.DBConn.Model(&models.StudentScore{}).
		Where(&models.StudentScore{SubjectCode: payload.SubjectCode, StudentCode: studentCode}).
		Find(&score)

	if score.StudentCode != "" && score.SubjectCode != "" {
		return c.Status(409).JSON(utils.ApiResponse(409, "Score is already exist", fiber.Map{}))
	}

	if studentCode != "" {
		newScore := models.StudentScore{
			StudentCode:          studentCode,
			SubjectCode:          payload.SubjectCode,
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

			database.DBConn.Model(&models.StudentScore{}).Create(&newScore)
			return c.Status(200).JSON(utils.ApiResponse(200, "Added score!", fiber.Map{}))
		}
	}

	return c.Status(400).JSON(utils.ApiResponse(400, "Your score is not valid", fiber.Map{}))
}

func GeneralScoresStatistics(c *fiber.Ctx) error {
	var numberOfStudents int64
	var numberOfDebtors int64
	var numberOfSubjects int64

	database.DBConn.Model(&models.StudentScore{}).Group("studentCode").Count(&numberOfStudents)
	database.DBConn.Model(&models.StudentScore{}).
		Where(&models.StudentScore{AlphabetScore: ""}).
		Where(&models.StudentScore{AlphabetScore: "F"}).
		Group("studentCode").
		Count(&numberOfDebtors)
	database.DBConn.Model(&models.SubjectInfo{}).Count(&numberOfSubjects)

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", fiber.Map{
		"numberOfStudents": numberOfStudents,
		"numberOfDebtors":  numberOfDebtors,
		"numberOfSubjects": numberOfSubjects,
	}))
}

func StudentStatistics(c *fiber.Ctx) error {
	type Result struct {
		Name                 string `gorm:"column:name" json:"name"`
		NumberOfCredit       int    `gorm:"column:noc" json:"numberOfCredit"`
		FirstComponentScore  string `gorm:"column:TP1" json:"firstComponentScore"`
		SecondComponentScore string `gorm:"column:TP2" json:"secondComponentScore"`
		ExamScore            string `gorm:"column:THI" json:"examScore"`
		AvgScore             string `gorm:"column:TONGKET" json:"avgScore"`
		AlphabetScore        string `gorm:"column:DIEMCHU" json:"alphabetScore"`
	}

	studentCode := strings.ToUpper(c.Params("studentCode"))
	var result []Result

	database.DBConn.Raw("SELECT * FROM studentScore inner join subjectInfo on studentScore.subjectCode = subjectInfo.subjectCode WHERE studentScore.studentCode = ?", studentCode).
		Find(&result)

	var failedSubjects []Result
	var passedSubjects []Result

	for _, score := range result {
		if !utils.IsPassedSubject(score.AlphabetScore) {
			failedSubjects = append(failedSubjects, score)
		}

		passedSubjects = append(passedSubjects, score)
	}

	var data fiber.Map

	if len(result) == 0 {
		data = fiber.Map{}
		return c.Status(404).JSON(utils.ApiResponse(404, "Student is not exist", data))
	}

	data = fiber.Map{
		"result":         result,
		"failedSubjects": failedSubjects,
		"passedSubjects": passedSubjects,
	}

	return c.Status(200).JSON(utils.ApiResponse(200, "OK", data))
}
