package models

import "gorm.io/gorm"

type AttendanceStudent struct {
	gorm.Model

	AttendanceID uint `json:"attendance_id"`
	Attendance   Attendance
	StudentID    uint `json:"student_id"`
	Student      Student
	Status       AttendanceStatus `json:"status"`
	Note         string           `json:"note"`
}
