package models

import (
	studentModel "api/internal/features/student/models"
	"gorm.io/gorm"
)

type SubjectAttendanceRecord struct {
	gorm.Model

	SubjectAttendanceId uint
	SubjectAttendance   SubjectAttendance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId           uint
	Student             studentModel.Student
}
