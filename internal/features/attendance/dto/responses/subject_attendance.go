package responses

import (
	"api/internal/features/attendance/domains"
	"api/internal/features/attendance/dto"
	user "api/internal/features/user/domains"
)

type GetAllSubjectAttendances struct {
	Items []dto.GetAllSubjectAttendancesItem `json:"items" validate:"required"`
} // @name GetAllSubjectAttendancesRes

type GetAllSubjectAttendanceRecords struct {
	Items []dto.GetAllSubjectAttendanceRecordsItem `json:"items" validate:"required"`
}

type GetSubjectAttendance struct {
	SubjectAttendance domains.SubjectAttendance `json:"subject_attendance" validate:"required"`
	Creator           user.User                 `json:"creator" validate:"required"`
} // @name GetSubjectAttendanceRes

type CreateSubjectAttendance struct {
	SubjectAttendance domains.SubjectAttendance `json:"subject_attendance" validate:"required"`
}

type CreateSubjectAttendanceRecord struct {
	SubjectAttendanceRecord domains.SubjectAttendanceRecord `json:"subject_attendance_record" validate:"required"`
}

type CreateSubjectAttendanceRecordStudent struct {
	Message string `json:"message" validate:"required"`
}

type DeleteSubjectAttendanceRecord struct {
	Message string `json:"message" validate:"required"`
}
