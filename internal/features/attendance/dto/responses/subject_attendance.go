package responses

import "api/internal/features/attendance/domains"

type GetAllSubjectAttendances struct {
	SubjectAttendances []domains.SubjectAttendance `json:"subject_attendances" validate:"required"`
}

type CreateSubjectAttendance struct {
	SubjectAttendance domains.SubjectAttendance `json:"subject_attendance" validate:"required"`
}
