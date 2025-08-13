package dto

import (
	"api/internal/features/attendance/domains"
	attendance "api/internal/features/attendance/domains"
	studentDomain "api/internal/features/student/domains"
	subject "api/internal/features/subject/domains"
)

type GetAllSubjectAttendanceRecordsItem struct {
	Student studentDomain.Student           `json:"student" validate:"required"`
	Record  domains.SubjectAttendanceRecord `json:"record" validate:"required"`
} // @name GetAllSubjectAttendanceRecordsItem

type GetAllSubjectAttendancesItem struct {
	SubjectAttendance attendance.SubjectAttendance `json:"subject_attendance" validate:"required"`
	Subject           subject.Subject              `json:"subject" validate:"required"`
} // @name GetAllSubjectAttendancesItem
