package models

import (
	models2 "api/internal/features/batch/models"
	"gorm.io/gorm"
)

type Major struct {
	gorm.Model

	Name    string        `gorm:"not null"`
	BatchId uint          `gorm:"not null"`
	Batch   models2.Batch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
