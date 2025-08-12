package dto

import (
	"api/internal/features/attendance/domains"
	studentDomain "api/internal/features/student/domains"
)

type GeneralAttendanceStudentItem struct {
	Student studentDomain.Student           `json:"student" validate:"required"`
	Record  domains.GeneralAttendanceRecord `json:"record" validate:"required"`
}
