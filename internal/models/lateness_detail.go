package models

import "gorm.io/gorm"

type LatenessDetail struct {
	gorm.Model

	LatenessId uint
	Lateness   Lateness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId  uint
	Student    Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
