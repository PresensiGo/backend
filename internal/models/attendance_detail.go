package models

import "gorm.io/gorm"

type AttendanceDetail struct {
	gorm.Model

	AttendanceId uint `json:"attendance_id"`
	Attendance   Attendance
	StudentId    uint `json:"student_id"`
	Student      Student
	Status       AttendanceStatus `json:"status"`
	Note         string           `json:"note"`
}
