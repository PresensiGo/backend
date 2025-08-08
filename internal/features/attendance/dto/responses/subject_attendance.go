package responses

import (
	"api/internal/features/attendance/domains"
	shared "api/internal/shared/domains"
)

type GetAllSubjectAttendances struct {
	Items []shared.SubjectAttendanceSubject `json:"items" validate:"required"`
}

type CreateSubjectAttendance struct {
	SubjectAttendance domains.SubjectAttendance `json:"subject_attendance" validate:"required"`
}
