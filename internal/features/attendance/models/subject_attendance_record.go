package models

import (
	"time"

	student "api/internal/features/student/models"
	"api/pkg/constants"
	"gorm.io/gorm"
)

type SubjectAttendanceRecord struct {
	gorm.Model

	DateTime            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	SubjectAttendanceId uint
	SubjectAttendance   SubjectAttendance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId           uint
	Student             student.Student
	Status              constants.AttendanceStatus
}
