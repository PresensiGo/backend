package models

import (
	models2 "api/internal/features/student/models"
	"gorm.io/gorm"
)

type LatenessDetail struct {
	gorm.Model

	LatenessId uint
	Lateness   Lateness `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId  uint
	Student    models2.Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
