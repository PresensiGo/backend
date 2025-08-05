package models

import (
	models2 "api/internal/features/classroom/models"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	NIS         string `gorm:"not null"`
	Name        string `gorm:"not null"`
	ClassroomId uint
	Classroom   models2.Classroom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
