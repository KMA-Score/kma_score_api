package models

type Score struct {
	ID                   int     `gorm:"column:Id" json:"-"`
	StudentId            string  `gorm:"column:StudentId" json:"-"`
	SubjectId            string  `gorm:"column:SubjectId" json:"-"`
	FirstComponentScore  string  `gorm:"column:FirstComponentScore" json:"firstComponentScore"`
	SecondComponentScore string  `gorm:"column:SecondComponentScore" json:"secondComponentScore"`
	ExamScore            string  `gorm:"column:ExamScore" json:"examScore"`
	AvgScore             string  `gorm:"column:AvgScore" json:"avgScore"`
	AlphabetScore        string  `gorm:"column:AlphabetScore" json:"alphabetScore"`
	Subject              Subject `json:"subject"`
}

func (Score) TableName() string {
	return "Scores"
}
