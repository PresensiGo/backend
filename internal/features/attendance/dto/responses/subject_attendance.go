package responses

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
)

type GetAllSubjectAttendances struct {
	Items []dto.GetAllSubjectAttendancesItem `json:"items" validate:"required"`
} // @name GetAllSubjectAttendancesRes

type GetAllSubjectAttendanceRecords struct {
	Items []dto.GetAllSubjectAttendanceRecordsItem `json:"items" validate:"required"`
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
