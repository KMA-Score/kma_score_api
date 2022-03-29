package models

type StudentScore struct {
	ID                   int    `gorm:"column:id" json:"id"`
	StudentCode          string `gorm:"column:studentCode" json:"studentCode"`
	SubjectCode          string `gorm:"column:subjectCode" json:"subjectCode"`
	FirstComponentScore  string `gorm:"column:TP1" json:"firstComponentScore"`
	SecondComponentScore string `gorm:"column:TP2" json:"secondComponentScore"`
	ExamScore            string `gorm:"column:THI" json:"examScore"`
	AvgScore             string `gorm:"column:TONGKET" json:"avgScore"`
	AlphabetScore        string `gorm:"column:DIEMCHU" json:"alphabetScore"`
}

func (StudentScore) TableName() string {
	return "studentScore"
}
