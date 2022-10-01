package models

type Subject struct {
	ID              string  `gorm:"column:Id" json:"id"`
	Name            string  `gorm:"column:Name" json:"name"`
	NumberOfCredits int     `gorm:"column:NumberOfCredits" json:"numberOfCredits"`
	Scores          []Score `gorm:"foreignKey:SubjectId" json:"-"`
}

func (Subject) TableName() string {
	return "Subjects"
}
