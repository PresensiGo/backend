package models

import (
	models2 "api/internal/features/student/models"
	"gorm.io/gorm"
)

type AttendanceDetail struct {
	gorm.Model

	AttendanceId uint             `json:"attendance_id"`
	Attendance   Attendance       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId    uint             `json:"student_id"`
	Student      models2.Student  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       AttendanceStatus `json:"status"`
	Note         string           `json:"note"`
}
