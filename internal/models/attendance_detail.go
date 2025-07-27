package models

import "gorm.io/gorm"

type AttendanceDetail struct {
	gorm.Model

	AttendanceId uint             `json:"attendance_id"`
	Attendance   Attendance       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StudentId    uint             `json:"student_id"`
	Student      Student          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       AttendanceStatus `json:"status"`
	Note         string           `json:"note"`
}
