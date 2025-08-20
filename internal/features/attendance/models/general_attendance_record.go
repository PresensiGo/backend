package models

import (
	"time"

	"api/internal/features/student/models"
	"gorm.io/gorm"
)

type GeneralAttendanceRecord struct {
	gorm.Model

	DateTime            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	GeneralAttendanceId uint
	GeneralAttendance   GeneralAttendance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId           uint
	Student             models.Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
