package models

type Student struct {
	ID     string  `gorm:"column:Id" json:"id"`
	Name   string  `gorm:"column:Name" json:"name"`
	Class  string  `gorm:"column:Class" json:"class"`
	Scores []Score `gorm:"foreignKey:StudentId" json:"scores"`
}

func (Student) TableName() string {
	return "Students"
}
