package dto

import (
	"api/internal/features/attendance/domains"
	studentDomain "api/internal/features/student/domains"
)

type SubjectAttendanceRecordItem struct {
	Student studentDomain.Student           `json:"student" validate:"required"`
	Record  domains.SubjectAttendanceRecord `json:"record" validate:"required"`
}
