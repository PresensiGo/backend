package models

import (
	"time"

	models2 "api/internal/features/classroom/models"
	"gorm.io/gorm"
)

type AttendanceStatus string

const (
	AttendancePresent    AttendanceStatus = "hadir"
	AttendancePermission AttendanceStatus = "izin"
	AttendanceSick       AttendanceStatus = "sakit"
	AttendanceAlpha      AttendanceStatus = "alpha"
)

// deprecated
type Attendance struct {
	gorm.Model

	ClassroomId uint              `json:"classroom_id"`
	Classroom   models2.Classroom `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date        time.Time         `json:"date"`
}
