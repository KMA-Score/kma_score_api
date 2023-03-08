package models

type ApiKey struct {
	ID     int    `gorm:"column:Id"`
	Key    string `gorm:"column:Key"`
	Secret string `gorm:"column:Secret"`
}

func (ApiKey) TableName() string {
	return "ApiKeys"
}
