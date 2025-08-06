package models

import (
	"api/internal/features/school/models"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model

	Name     string
	SchoolId uint
	School   models.School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
