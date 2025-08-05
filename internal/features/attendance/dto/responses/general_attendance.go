package responses

import "api/internal/features/attendance/domains"

type CreateGeneralAttendance struct {
	GeneralAttendance domains.GeneralAttendance `json:"general_attendance" validate:"required"`
}

type GetAllGeneralAttendances struct {
	GeneralAttendances []domains.GeneralAttendance `json:"general_attendances" validate:"required"`
}
