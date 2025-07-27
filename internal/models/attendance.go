package models

import (
	"gorm.io/gorm"
	"time"
)

type AttendanceStatus string

const (
	AttendancePresent    AttendanceStatus = "hadir"
	AttendancePermission AttendanceStatus = "izin"
	AttendanceSick       AttendanceStatus = "sakit"
	AttendanceAlpha      AttendanceStatus = "alpha"
)

type Attendance struct {
	gorm.Model

	ClassroomId uint      `json:"classroom_id"`
	Classroom   Classroom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date        time.Time `json:"date"`
}
