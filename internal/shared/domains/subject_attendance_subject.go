package domains

import (
	attendance "api/internal/features/attendance/domains"
	subject "api/internal/features/subject/domains"
)

type SubjectAttendanceSubject struct {
	SubjectAttendance attendance.SubjectAttendance `json:"subject_attendance" validate:"required"`
	Subject           subject.Subject              `json:"subject" validate:"required"`
}
