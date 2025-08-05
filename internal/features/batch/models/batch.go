package models

import (
	models2 "api/internal/features/school/models"
	"gorm.io/gorm"
)

type Batch struct {
	gorm.Model

	Name     string `gorm:"not null"`
	SchoolId uint
	School   models2.School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
