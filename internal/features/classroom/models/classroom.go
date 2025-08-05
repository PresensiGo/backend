package models

import (
	models2 "api/internal/features/major/models"
	"gorm.io/gorm"
)

type Classroom struct {
	gorm.Model

	Name    string        `gorm:"not null"`
	MajorId uint          `gorm:"not null"`
	Major   models2.Major `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
