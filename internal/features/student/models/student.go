package models

import (
	classroom "api/internal/features/classroom/models"
	school "api/internal/features/school/models"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	NIS         string `gorm:"not null"`
	Name        string `gorm:"not null"`
	SchoolId    uint
	School      school.School `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ClassroomId uint
	Classroom   classroom.Classroom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Gender      string
}
