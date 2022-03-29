package utils

import "kma_score_api/models"

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

func GetSubjectIndex(subjects []models.SubjectInfo, element models.StudentScore) int {
	for i, e := range subjects {
		if e.SubjectCode == element.SubjectCode {
			return i
		}
	}

	return -1
}
