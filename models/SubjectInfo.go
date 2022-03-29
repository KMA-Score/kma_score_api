package models

type SubjectInfo struct {
	ID             int    `gorm:"column:id" json:"id"`
	SubjectCode    string `gorm:"column:subjectCode" json:"subjectCode"`
	Name           string `gorm:"column:name" json:"name"`
	NumberOfCredit int    `gorm:"column:noc" json:"numberOfCredit"`
}

func (SubjectInfo) TableName() string {
	return "subjectInfo"
}
