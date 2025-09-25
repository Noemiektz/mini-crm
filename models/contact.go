package models

type Contact struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Nom   string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;unique;not null"`
}
