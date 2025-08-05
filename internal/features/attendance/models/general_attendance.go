package models

import (
	"time"

	"api/internal/features/school/models"
	"gorm.io/gorm"
)

type GeneralAttendance struct {
	gorm.Model

	Date     time.Time
	DueTime  time.Time
	Note     string
	SchoolId uint
	School   models.School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
