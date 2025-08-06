package models

import (
	"time"

	classroom "api/internal/features/classroom/models"
	subject "api/internal/features/subject/models"
	"gorm.io/gorm"
)

type SubjectAttendance struct {
	gorm.Model

	DateTime    time.Time
	Code        string
	Note        string
	ClassroomId uint
	Classroom   classroom.Classroom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubjectId   uint
	Subject     subject.Subject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
