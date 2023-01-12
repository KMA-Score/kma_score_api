package utils

import (
	"fmt"
	"kma_score_api/models"
	"regexp"
	"strconv"
)

func AlphabetScoreToTetraScore(alphabetScore string) float64 {
	switch alphabetScore {
	case "A+":
		return 4.0
	case "A":
		return 3.8
	case "B+":
		return 3.5
	case "B":
		return 3.0
	case "C+":
		return 2.5
	case "C":
		return 2
	case "D+":
		return 1.5
	case "D":
		return 1
	case "F":
		return 0
	default:
		return 0
	}
}

func GetSubjectIndex(subjects []models.Subject, element models.Score) int {
	for i, e := range subjects {
		if e.ID == element.SubjectId {
			return i
		}
	}

	return -1
}

func IsValidScore(score models.Score) bool {
	firstComponentScore, err := strconv.ParseFloat(score.FirstComponentScore, 64)
	secondComponentScore, err := strconv.ParseFloat(score.SecondComponentScore, 64)
	examScore, err := strconv.ParseFloat(score.ExamScore, 64)
	avgScore, err := CalcSubjectAvgScore(score)

	if score.AlphabetScore == "" && err != nil {
		return false
	}

	if firstComponentScore < 0 || firstComponentScore > 10 || secondComponentScore < 0 || secondComponentScore > 10 || examScore < 0 || examScore > 10 {
		return false
	}

	if score.AlphabetScore != "" && (ConvertDecimalScoreToAlphabetScore(avgScore) != score.AlphabetScore) {
		return false
	}

	return true
}

func CalcSubjectAvgScore(score models.Score) (string, error) {
	firstComponentScore, err := strconv.ParseFloat(score.FirstComponentScore, 64)
	secondComponentScore, err := strconv.ParseFloat(score.SecondComponentScore, 64)
	examScore, err := strconv.ParseFloat(score.ExamScore, 64)

	avgScore := (firstComponentScore*0.3+secondComponentScore*0.7)*0.3 + examScore*0.7

	return fmt.Sprintf("%.1f", avgScore), err
}

func ConvertDecimalScoreToAlphabetScore(stringDecimalScore string) string {
	alphabetScore := "F"
	decimalScore, _ := strconv.ParseFloat(stringDecimalScore, 64)

	if decimalScore <= 4 {
		alphabetScore = "F"
	} else if decimalScore <= 4.7 {
		alphabetScore = "D"
	} else if decimalScore <= 5.4 {
		alphabetScore = "D+"
	} else if decimalScore <= 6.2 {
		alphabetScore = "C"
	} else if decimalScore <= 6.9 {
		alphabetScore = "C+"
	} else if decimalScore <= 7.7 {
		alphabetScore = "B"
	} else if decimalScore <= 8.4 {
		alphabetScore = "B+"
	} else if decimalScore <= 8.9 {
		alphabetScore = "A"
	} else if decimalScore <= 10 {
		alphabetScore = "A+"
	}

	return alphabetScore
}

func IsPassedSubject(alphabetScore string) bool {
	if alphabetScore == "" || alphabetScore == "F" {
		return false
	}

	return true
}

func ShouldCalculateAverageScore(score models.Score) bool {
	matched, _ := regexp.Match(`ATQGTC\d+`, []byte(score.SubjectId))

	if matched {
		return false
	}

	return true
}
