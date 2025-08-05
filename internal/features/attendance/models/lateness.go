package models

import (
	"time"

	models2 "api/internal/features/school/models"
	"gorm.io/gorm"
)

type Lateness struct {
	gorm.Model

	Date     time.Time
	SchoolId uint
	School   models2.School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
