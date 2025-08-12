package responses

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
	shared "api/internal/shared/domains"
)

type GetAllSubjectAttendances struct {
	Items []shared.SubjectAttendanceSubject `json:"items" validate:"required"`
}

type GetAllSubjectAttendanceRecords struct {
	Items []dto.SubjectAttendanceRecordItem `json:"items" validate:"required"`
}

type GetSubjectAttendance struct {
	SubjectAttendance domains.SubjectAttendance `json:"subject_attendance" validate:"required"`
}

type CreateSubjectAttendance struct {
	SubjectAttendance domains.SubjectAttendance `json:"subject_attendance" validate:"required"`
}

type CreateSubjectAttendanceRecordStudent struct {
	Message string `json:"message" validate:"required"`
}
