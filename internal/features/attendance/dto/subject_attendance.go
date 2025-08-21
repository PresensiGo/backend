package dto

import (
	"api/internal/features/attendance/domains"
	attendance "api/internal/features/attendance/domains"
	studentDomain "api/internal/features/student/domains"
	subject "api/internal/features/subject/domains"
	user "api/internal/features/user/domains"
)

type GetAllSubjectAttendanceRecordsItem struct {
	Student studentDomain.Student           `json:"student" validate:"required"`
	Record  domains.SubjectAttendanceRecord `json:"record" validate:"required"`
} // @name GetAllSubjectAttendanceRecordsItem

type GetAllSubjectAttendancesItem struct {
	SubjectAttendance attendance.SubjectAttendance `json:"subject_attendance" validate:"required"`
	Subject           subject.Subject              `json:"subject" validate:"required"`
	Creator           user.User                    `json:"creator" validate:"required"`
} // @name GetAllSubjectAttendancesItem

type GetAllSubjectAttendancesStudentItem struct {
	SubjectAttendance       attendance.SubjectAttendance       `json:"subject_attendance" validate:"required"`
	SubjectAttendanceRecord attendance.SubjectAttendanceRecord `json:"subject_attendance_record" validate:"required"`
	Subject                 subject.Subject                    `json:"subject" validate:"required"`
	Creator                 user.User                          `json:"creator" validate:"required"`
} // @name GetAllSubjectAttendancesStudentItem
