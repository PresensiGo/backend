package models

import (
	"api/internal/features/student/models"
	"gorm.io/gorm"
)

type GeneralAttendanceRecord struct {
	gorm.Model

	GeneralAttendanceId uint
	GeneralAttendance   GeneralAttendance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId           uint
	Student             models.Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
